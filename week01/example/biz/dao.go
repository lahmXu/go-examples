package biz

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type User struct {
	ID   int
	Name string
}

func GetOneEntity() (User, error) {
	var res User
	result := db.Where("id = 3").Find(&res)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, errors.Wrap(result.Error, "user: record not found")
		}
		return User{}, result.Error
	}
	return res, nil
}
