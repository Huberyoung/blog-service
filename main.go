package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"blog-service/global"
	"blog-service/interbal/model"
	"blog-service/interbal/routers"
	"blog-service/pkg/setting"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	if err := setUpSetting(); err != nil {
		log.Fatalf("init.setUpSetting err:%v\n", err)
	}

	if err := setUpDBEngine(); err != nil {
		log.Fatalf("init.setUpDBEngine err:%v\n", err)
	}
}

func main() {
	route := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           route,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    global.ServerSetting.MaxHeaderBytes,
	}
	err := s.ListenAndServe()
	fmt.Printf("err:%s\n", err.Error())
}

func setUpSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = set.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("Database", &global.DataBaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.ServerSetting.MaxHeaderBytes = 1 << global.ServerSetting.MaxHeaderBytes
	return nil
}

func setUpDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDbEngine(global.DataBaseSetting)
	return err
}
