package v1

import (
	"github.com/gin-gonic/gin"
)

// User 用户
type User struct {
	ID   int
	Name string
}

// IUser IUser
type IUser interface {
	GetUserInfo(c *gin.Context)
}
