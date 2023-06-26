package control

import (
	"go-vue/common"
	"go-vue/model"
	"go-vue/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Rrgister(c *gin.Context) {
	db := common.GetDB()
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
	if util.ExistTelephone(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户存在"})
		return
	}
	//名称没有传则随机
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//创建用户
	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&user)

	//返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
