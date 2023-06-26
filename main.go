package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/locales/root"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null,unique"`
	Password  string `gotm:"varchar(255);not null"`
}

func main() {
	db := InitDB()
	fmt.Println("begin")
	r := gin.Default()
	r.GET("/api/auth/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		//数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号为11位"})
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不小于6位"})
		}
		if ExistTelephone(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户存在"})
			return
		}
		//名称没有传则随机
		if len(name) == 0 {
			name = RandomString(10)
		}

		//创建用户
		user := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&user)

		//返回结果
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnm")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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
	db.AutoMigrate(&User{})
	return db
}

func ExistTelephone(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
