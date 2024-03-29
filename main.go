package main

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/tracer"
	"context"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	if err := setUpLogger(); err != nil {
		log.Fatalf("init.setUpLogger err:%v\n", err)
	}

	if err := setUpTracing(); err != nil {
		log.Fatalf("init.setUpTracing err:%v\n", err)
	}
}

// @title			GO博客系统
// @version		1.0
// @description	练手项目，博客系统
// @termsOfService	https://github.com/Huberyoung/blog-service
// @host			localhost:8080
// @BasePath		/api/v1
// @contact.name	博客系统
// @contact.url	https://github.com/Huberyoung/blog-service
// @contact.email	huberyoung@163.com
func main() {
	route := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           route,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    global.ServerSetting.MaxHeaderBytes,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err:%s", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ...")

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeout); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	fmt.Println("Server exiting")
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
	global.AppSetting.UploadImageMaxSize *= 1024 * 1024
	global.AppSetting.DefaultContextTimeout *= time.Second

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

	err = set.ReadSection("JWT", &global.JwtSetting)
	if err != nil {
		return err
	}
	global.JwtSetting.Expire *= time.Second

	err = set.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	return nil
}

func setUpDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDbEngine(global.DataBaseSetting)
	return err
}

func setUpLogger() error {
	writer := &lumberjack.Logger{
		Filename:   global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:    global.AppSetting.LogMaxSize,
		MaxAge:     global.AppSetting.LogMaxAge,
		MaxBackups: global.AppSetting.LogMaxBackups,
		LocalTime:  global.AppSetting.LogUseLocalTime,
		Compress:   global.AppSetting.LogCompress,
	}
	global.Logger = logger.NewLogger(writer, "", log.LstdFlags).CloneWithCaller(2)
	return nil
}

func setUpTracing() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
