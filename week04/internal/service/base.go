package service

import (
	"github.com/gin-gonic/gin"
	errorCode "go-examples/week04/api/errors"
)

type Gin struct {
	Context *gin.Context
}

func (c *Gin) Response(httpCode int, data interface{}) {
	c.Context.JSON(httpCode, errorCode.Response{
		Data: data,
	})
	return
}
