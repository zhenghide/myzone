package controller

import (
	"github.com/gin-gonic/gin"
	"myzone/utils"
	"net/http"
	"myzone/model"
	"myzone/db"
)

func GetDiaries(c *gin.Context){
	user := utils.GetSession(c,"user")
	if user == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": "请登录！",
		})
	}else {
		diaries,_ := QueryDiariesByUserName((user).(*model.User).UserName)
		c.HTML(http.StatusOK, "diary.html", gin.H{
			"msg": "Diary",
			"user": user,
			"diaries": diaries,
		})
	}
}

//查询日志
func QueryDiariesByUserName(name string) ([]model.Diary, error){
	diaries := []model.Diary{}
	if rdb := db.DB().Order("create_time desc").Where("author = ?",name).Find(&diaries); rdb.Error != nil {
		return nil, rdb.Error
	}else {
		return diaries,nil
	}
}

//发表日志（新增）
func CreateDiary(c *gin.Context)  {
	user := utils.GetSession(c,"user")
	u,_ := (user).(*model.User)
	title := c.DefaultPostForm("diaryTitle","")
	content := c.DefaultPostForm("diaryContent","")
	createTime := utils.GetNowTimeLong()
	diary := model.Diary{
		Id:utils.GetUUID(),
		Title:title,
		Author:u.UserName,
		Content:content,
		CreateTime:createTime,
	}
	if rdb := db.DB().Create(&diary); rdb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "写日志失败!",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"mag": "发表成功!",
		})
	}
}


//删除日志
func DeleteDiary(c *gin.Context)  {
	diaryId := c.DefaultPostForm("diaryId","")
	diary := model.Diary{}
	if rdb := db.DB().Where("id = ?",diaryId).Delete(&diary); rdb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "删除失败!",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "发表成功!",
		})
	}
}
