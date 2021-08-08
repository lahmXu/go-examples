package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	config "go-examples/week04/configs"
	db "go-examples/week04/internal/respository"
	"go-examples/week04/internal/service"
	"net/http"
)

func main() {
	db.Init()
	serverInit()

}

func serverInit() {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	// 配置日志输出对象为zap logger
	g.Use(gin.Recovery())

	r := g.Group("/api")
	{
		version := r.Group("/user")
		version.GET("/:id", service.UserApi.GetUserInfo)

	}
	// start http server
	listenAddress := fmt.Sprintf(":%d", config.YamlConfig.App.Port)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           listenAddress,
		Handler:        g,
		ReadTimeout:    config.YamlConfig.App.ReadTimeout,
		WriteTimeout:   config.YamlConfig.App.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	server.ListenAndServe()

}
