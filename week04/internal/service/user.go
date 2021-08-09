package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	v1 "go-examples/week04/api/myapp/user/v1"
	"go-examples/week04/internal/domain"
	"strconv"
)

type UserService struct {
	domain.UserDomain
}

func (u UserService) GetUserInfo(c *gin.Context) {
	ctx := Gin{Context: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctx.Response(400, errors.New("param error."))
		return
	}
	result, err := u.GetUserByID(id)
	if err != nil {
		ctx.Response(500, errors.New("server error."))
		return
	}
	user := v1.User{}
	if err = copier.Copy(&user, &result); err != nil {
		ctx.Response(500, errors.New("server error."))
		return
	}
	ctx.Response(200, user)
	return
}
