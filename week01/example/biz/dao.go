package biz

import (
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
		return res, errors.Wrap(result.Error, "query error")
	}
	return res, nil
}
