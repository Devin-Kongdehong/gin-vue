package common

import (
	"fmt"
	"go-vue/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "go-vue"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%stcp(%s:%s)/%s?charset=%sparseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
