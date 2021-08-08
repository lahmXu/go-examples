package v1

import "github.com/gin-gonic/gin"

// User 用户
type User struct {
	ID   int
	Name string
}

// UserApi UserApi
type UserApi interface {
	GetUserInfo(c *gin.Context)
}
