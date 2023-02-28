package main

import (
	"blog-service/interbal/routers"
	"fmt"
	"net/http"
	"time"
)

func main() {
	route := routers.NewRouter()

	s := &http.Server{
		Addr:              ":8080",
		Handler:           route,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	err := s.ListenAndServe()
	fmt.Printf("err:%s\n", err.Error())
}
