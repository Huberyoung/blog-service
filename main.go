package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"service/global"
	"service/internal/model"
	"service/internal/routers"
	"service/pkg/logger"
	"service/pkg/setting"
	"time"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}

	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}

	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	global.Logger.InfoF("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")

	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		ReadTimeout:       global.ServerSetting.ReadTimeOut,
		ReadHeaderTimeout: global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDbEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
