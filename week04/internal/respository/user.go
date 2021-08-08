package respository

import (
	"go-examples/week04/internal/domain"
)

type UserDomain struct {
}

// GetUserByID 根据id查询单个用户
func (d UserDomain) GetUserByID(id int) (*domain.User, error) {
	var User domain.User
	err := db.Where("id = ?", id).Find(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}
