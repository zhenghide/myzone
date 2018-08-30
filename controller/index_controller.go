package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myzone/utils"
)

func GotoIndex(c *gin.Context){
	user := utils.GetSession(c,"user")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "Index",
		"user": user,
	})
}

func GotoLogin(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func GotoRegister(c *gin.Context){
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

