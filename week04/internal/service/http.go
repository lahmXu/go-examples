package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-examples/week04/api/myapp/user/v1"
	config "go-examples/week04/configs"
	"net/http"
)

type biz struct {
	userBiz v1.IUser
}

func Init() error {
	var biz biz
	return biz.ServerInit()
}

func (biz *biz) ServerInit() error {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	// 配置日志输出对象为zap logger
	g.Use(gin.Recovery())

	r := g.Group("/api")
	{
		userGroup := r.Group("/user")

		userGroup.GET("/:id", biz.userBiz.GetUserInfo)

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
	return server.ListenAndServe()

}
