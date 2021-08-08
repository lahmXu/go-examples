package biz

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"log"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	fmt.Println("========================")
	_, resErr := GetOneEntity()
	if resErr != nil {
		if errors.Is(resErr, gorm.ErrRecordNotFound) {
			fmt.Println("result is empty")
		} else {
			fmt.Printf("%+v", resErr)
		}
	}
}
