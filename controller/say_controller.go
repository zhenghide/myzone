package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myzone/model"
	"myzone/db"
	"myzone/utils"
)

func GetSays(c *gin.Context){
	user := utils.GetSession(c,"user")
	if user == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": "请登录！",
		})
	}else {
		says,_ := QuerySaysByUserName((user).(*model.User).UserName)
		c.HTML(http.StatusOK, "say.html", gin.H{
			"msg": "Say",
			"user": user,
			"says": says,
		})
	}
}

//发表说说（新增）
func CreateSay(c *gin.Context)  {
	user := utils.GetSession(c,"user")
	u,_ := (user).(*model.User)
	content := c.DefaultPostForm("sayContent","")
	createTime := utils.GetNowTimeLong()
	say := model.Say{
		Id:utils.GetUUID(),
		Author:u.UserName,
		Content:content,
		CreateTime:createTime,
	}
	if rdb := db.DB().Create(&say); rdb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "说说发表失败!",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"mag": "发表成功!",
		})
	}
}


//删除说说
func DeleteSay(c *gin.Context)  {
	sayId := c.DefaultPostForm("sayId","")
	say := model.Say{}
	if rdb := db.DB().Where("id = ?",sayId).Delete(&say); rdb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "删除失败!",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "发表成功!",
		})
	}
}

//查询说说
func QuerySaysByUserName(name string) ([]model.Say, error){
	says := []model.Say{}
	if rdb := db.DB().Order("create_time desc").Where("author = ?",name).Find(&says); rdb.Error != nil {
		return nil, rdb.Error
	}else {
		return says,nil
	}
}