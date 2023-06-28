package main

import (
	"fmt"
	"go-vue/common"
	"go-vue/control"
	"go-vue/middleware"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	//"github.com/go-playground/locales/root"
	//"github.com/jinzhu/gorm"
)

func main() {
	InitConfig()
	db := common.InitDB()
	fmt.Println(db.Config)
	r := gin.Default()
	r.POST("/api/auth/register", control.Rrgister)
	r.POST("/api/auth/login", control.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), control.Info)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
