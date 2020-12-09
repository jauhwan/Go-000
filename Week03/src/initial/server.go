package initial

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func InitHttpServer(ctx context.Context, exit <-chan os.Signal) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
	})
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: http.DefaultServeMux,
	}
	go func() {
		<-exit
		// 一分钟超时时间等待优雅结束server
		ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal("httpserver shutdown error:%v", err)
		}
	}()
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
