package common

import (
	"fmt"
	"go-vue/model"

	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	fmt.Print(1)
	host := "localhost"
	port := "3306"
	database := "mysql"
	username := "root"
	password := "123456"
	charset := "utf8mb4"
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	//gorm.Open(sqlite.Open(args), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
