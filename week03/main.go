package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建 context. Background()主要用于main函数、初始化以及测试代码中
	g, ctx := errgroup.WithContext(context.Background())

	// 创建 server,增加配置信息
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world!"))
	})
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// 创建 channel 用来传输错误信息
	serverExit := make(chan int)

	mux.HandleFunc("/error", func(writer http.ResponseWriter, request *http.Request) {
		serverExit <- 500
	})
	mux.HandleFunc("/done", func(writer http.ResponseWriter, request *http.Request) {
		server.Shutdown(ctx)
	})

	// 通过 goroutine 启动服务
	g.Go(func() error {
		fmt.Printf("http server start %s", server.Addr)
		return server.ListenAndServe()

	})
	// 监听 channel,停止 server 服务
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("g2: context done. exiting...")
		case <-serverExit:
			log.Println("g2: server error. exiting...")
		}
		return server.Shutdown(ctx)
	})

	g.Go(func() error {
		// 监听信号, ctrl+c 是发送 SIGINT 信号，终止一个进程
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Println("g3: context done. exiting...")
			return ctx.Err()
		case sig := <-quit:
			log.Printf("g3: receive os signal: %v", sig)
			return errors.New("receive quit signal.")
		}
	})

	log.Printf("[main] errgroup exiting: %+v\n", g.Wait())
}
