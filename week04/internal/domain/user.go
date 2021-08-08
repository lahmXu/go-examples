package domain

import "go-examples/week04/internal/respository"

// User 用户实体
type User struct {
	respository.Model
	Name string `gorm:"column:name;not null;type:varchar(50);commnet:'名称'" json:"name"` //名称
}

type UserDomain interface {
	GetUserByID(id int) (*User, error)
}
