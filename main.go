package main

import (
	"fmt"
	"go-vue/common"
	"go-vue/control"
	"go-vue/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-playground/locales/root"
	//"github.com/jinzhu/gorm"
)

func main() {
	db := common.InitDB()
	fmt.Println(db.Config)
	r := gin.Default()
	r.POST("/api/auth/register", control.Rrgister)
	r.POST("/api/auth/login", control.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), control.Info)
	r.Run()
}
