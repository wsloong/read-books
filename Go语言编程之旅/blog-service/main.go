package main

import (
	"log"
	"net/http"
	"path"
	"time"

	"github.com/wsloong/blog-service/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/wsloong/blog-service/global"
	"github.com/wsloong/blog-service/internal/model"
	"github.com/wsloong/blog-service/internal/routers"
	setting2 "github.com/wsloong/blog-service/pkg/setting"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger error: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 编程之旅： 一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
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

// 初始化配置
func setupSetting() error {
	setting, err := setting2.NewSetting()
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

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
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
