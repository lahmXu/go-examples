package respository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "go-examples/week04/configs"
	"log"
	"time"
)

var db *gorm.DB

// Model 基础表结构
type Model struct {
	ID        int       `gorm:"column:id;primaryKey;type:int(11)" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;type:datetime;commnet:'创建时间'" json:"created_at"`                       //创建时间
	CreatedBy int       `gorm:"column:created_by;not null;type:int(11);commnet:'创建人'" json:"created_by"`                         //创建人
	UpdatedAt time.Time `gorm:"column:updated_at;not null;type:datetime;commnet:'更新时间'" json:"updated_at"`                       //更新时间
	UpdatedBy int       `gorm:"column:updated_by;not null;type:int(11);commnet:'更新人'" json:"updated_by"`                         //更新人
	IsDeleted int8      `gorm:"column:is_deleted;not null;default:0;type:tinyint(4);commnet:'是否删除，0-存在，1-删除'" json:"is_deleted"` //是否删除，0-存在，1-删除
}

func Init() error {
	var err error
	db, err = gorm.Open(config.YamlConfig.DB.DbType, config.YamlConfig.DB.ConnStr)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
		return err
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return nil
}
