package router

import (
	"github.com/gin-gonic/gin"
	"myzone/controller"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)

func InitRouter() *gin.Engine{
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	//设置静态资源路径
	router.Static("/static","./static")

	router.LoadHTMLGlob("view/*")

	//首页
	router.GET("/index", controller.GotoIndex)

	//跳转登录页面
	router.GET("/page/login", controller.GotoLogin)

	//登陆
	router.POST("/login", controller.Login)

	//注销
	router.GET("/logout", controller.Logout)

	//跳转注册页面
	router.GET("/page/register", controller.GotoRegister)

	//注册(创建用户)
	router.POST("/register", controller.CreateUser)

	//跳转个人主页
	router.GET("/home", controller.GotoHome)

	//跳转说说页面
	router.GET("/say", controller.GetSays)

	//发表说说
	router.POST("/say/create", controller.CreateSay)

	//删除说说
	router.POST("/say/delete", controller.DeleteSay)

	//跳转日志页面
	router.GET("/diary", controller.GetDiaries)

	//写日志
	router.POST("/diary/create", controller.CreateDiary)

	//删除日志
	router.POST("/diary/delete", controller.DeleteDiary)

	//跳转相册页面
	router.GET("/photo", controller.GetPhotos)

	//创建相册
	router.POST("/photo/create", controller.CreatePhoto)

	//上传照片
	router.POST("/picture/upload", controller.UploadPicture)

	return router
}
