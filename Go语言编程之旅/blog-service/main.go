package main

import (
	"log"
	"net/http"
	"time"

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
		log.Fatalf("init.setupDBEngine err:%v", err)
	}
}

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
