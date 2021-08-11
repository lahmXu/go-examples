package main

import (
	"context"
	"github.com/pkg/errors"
	db "go-examples/week04/internal/respository"
	"go-examples/week04/internal/service"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return db.Init()
	})

	g.Go(func() error {
		return service.Init()
	})

	g.Go(func() error {
		// 监听信号, ctrl+c 是发送 SIGINT 信号，终止一个进程
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Println("context done. exiting...")
			return ctx.Err()
		case sig := <-quit:
			log.Printf("receive os signal: %v", sig)
			return errors.New("receive quit signal.")
		}
	})

	log.Printf("[main] exiting: %+v\n", g.Wait())
}
