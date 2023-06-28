package control

import (
	"go-vue/common"
	"go-vue/dto"
	"go-vue/model"
	"go-vue/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
	}

	//创建用户
	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	db.Create(&user)

	//返回结果
	c.JSON(200, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号为11位"})
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不小于6位"})
	}

	//判断用户是否存在和密码是否正确
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码不正确"})
		return
	}

	//返回
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "token生成失败"})
		log.Printf("token err:%v", err)
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"message": "注册成功",
	})

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},//控制不输出敏感信息
	})
}
