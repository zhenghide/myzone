package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	"myzone/model"
	"myzone/db"
	"myzone/utils"
)

//登陆
func Login(c *gin.Context){
	userName := c.DefaultPostForm("userName","")
	password := c.DefaultPostForm("password","")
	user,_ := CheckUser(userName, password)
	if user != nil {
		//设置session
		utils.SetSession(c,"user", user)

		c.JSON(http.StatusOK, gin.H{
			"msg": "正在登陆！",
			"user": user,
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"error": "用户名或密码错误!",
		})
	}
}

//注销
func Logout(c *gin.Context){
	//删除session
	utils.DeleteSession(c,"user")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "注销成功！",
	})
}

func CheckUser(name string, pwd string) (*model.User, error){
	user := model.User{
		UserName: name,
		Password: pwd,
	}

	var rdb *gorm.DB
	rdb = db.DB().Where("user_name = ? AND password = ?", name, pwd).Find(&user)
	if rdb.Error != nil {
		return nil,rdb.Error
	}
	return &user,nil
}

func QueryByName(name string) (*model.User, error){
	user := model.User{
		UserName: name,
	}

	var rdb *gorm.DB
	rdb = db.DB().Where("user_name = ?", name).Find(&user)
	if rdb.Error != nil {
		return nil,rdb.Error
	}
	return &user,nil
}

func CreateUser(c *gin.Context){
	nickName := c.DefaultPostForm("nickName","")
	userName := c.DefaultPostForm("userName","")
	password := c.DefaultPostForm("password","")
	phone := c.DefaultPostForm("phone","")
	email := c.DefaultPostForm("email","")
	user := model.User{
		Id:utils.GetUUID(),
		NickName:nickName,
		UserName:userName,
		Password:password,
		Phone:phone,
		Email:email,
	}

	u,_ := QueryByName(userName)
	if u !=nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "用户已存在!",
		})
	}else {
		var rdb *gorm.DB
		rdb = db.DB().Create(&user)
		if rdb.Error !=nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "注册失败！",
			})
		}else {
			utils.SetSession(c,"user", user)
			c.JSON(http.StatusOK, gin.H{
				"msg": "注册成功！",
			})
		}

	}
}



