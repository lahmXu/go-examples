package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-examples/week04/api/myapp/user/v1"
	config "go-examples/week04/configs"
	db "go-examples/week04/internal/respository"
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
		userGroup := r.Group("/user")

		// TODO 这边调用这个接口(v1.IUser.GetUserInfo)一直报错,求指导
		userGroup.GET("/:id", v1.IUser.GetUserInfo)

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
