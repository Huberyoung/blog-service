package main

import (
	"blog-service/global"
	"blog-service/interbal/routers"
	"blog-service/pkg/setting"
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
	if err := setUpSetting(); err != nil {
		log.Fatalf("init.setUpSetting err:%v\n", err)
	}
}

func main() {
	route := routers.NewRouter()

	fmt.Printf("ServerSetting:%v\n", global.ServerSetting)
	fmt.Printf("AppSetting:%v\n", global.AppSetting)
	fmt.Printf("DataBaseSetting:%v\n", global.DataBaseSetting)

	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           route,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
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
	return nil
}
