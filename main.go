package main

import (
	"fmt"
	"go-vue/common"
	"go-vue/control"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/locales/root"
	//"github.com/jinzhu/gorm"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	fmt.Println("begin")
	r := gin.Default()
	r.GET("/api/auth/register", control.Rrgister)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
