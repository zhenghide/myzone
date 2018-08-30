package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myzone/utils"
)

func GotoHome(c *gin.Context){
	user := utils.GetSession(c,"user")
	if user == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": "请登录！",
		})
	}else {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"msg": "Home",
			"user": user,
		})
	}
}
