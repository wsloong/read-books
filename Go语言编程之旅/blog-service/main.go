package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/wsloong/blog-service/pkg/tracer"

	"github.com/wsloong/blog-service/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/wsloong/blog-service/global"
	"github.com/wsloong/blog-service/internal/model"
	"github.com/wsloong/blog-service/internal/routers"
	setting2 "github.com/wsloong/blog-service/pkg/setting"
)

var (
	port    string
	runMode string
	config  string

	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	if err := setupFlag(); err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}

	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	if err := setupTracer(); err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 编程之旅： 一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// 读取命令行参数
func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()
	return nil
}

// 初始化配置
func setupSetting() error {
	setting, err := setting2.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}

	if err = setting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("Email", &global.EmailSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("JWT", &global.JWTSetting); err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

// 初始化数据库
func setupDBEngine() (err error) {
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}

// 初始化日志模块
func setupLogger() error {
	// 日志文件路径
	logFile := path.Join(global.AppSetting.LogSavePath, global.AppSetting.LogFileName+global.AppSetting.LogFileExt)

	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  logFile,
			MaxSize:   600,  // 日志文件最大为600M
			MaxAge:    10,   // 日志文件最大生成周期为10天
			LocalTime: true, // 日志文件名的时间格式为本地时间
		},
		"", log.LstdFlags).WithCaller(2)
	return nil
}

// 初始化链路追踪模块
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return nil
	}
	global.Tracer = jaegerTracer
	return nil
}
