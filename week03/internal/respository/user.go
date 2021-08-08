package respository

import "github.com/jinzhu/gorm"

// User 用户实体
type User struct {
	Model
	Name string `gorm:"column:name;not null;type:varchar(50);commnet:'名称'" json:"name"` //名称
}

// GetUserByID 根据id查询单个用户
func GetUserByID(id int) (*User, error) {
	var User User
	err := db.Where("id = ?", id).Find(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

// ListUser 查询用户列表
func ListUser(maps interface{}) (result []*User, err error) {
	err = db.Where(maps).Find(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return result, nil
}

// RemoveUser 移除用户
func RemoveUser(id int) (int64, error) {
	l, err := GetUserByID(id)
	if l == nil || err != nil {
		return 0, err
	}
	res := db.Model(&User{}).Where("id = ?", id).Update("is_deleted", 1)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

// AddUser 添加用户信息
func AddUser(User User) error {
	if err := db.Create(&User).Error; err != nil {
		return err
	}
	return nil
}

// EditUser 编辑用户信息
func EditUser(id int, user User) (int64, error) {
	l, err := GetUserByID(id)
	if l == nil || err != nil {
		return 0, err
	}
	result := db.Model(&User{}).Where("id = ? and is_deleted = ?", id, 0).Update(user)
	if result.Error != nil {
		return result.RowsAffected, err
	}
	return result.RowsAffected, nil
}
