package controller

import (
	"github.com/gin-gonic/gin"
	"myzone/utils"
	"net/http"
	"myzone/model"
	"myzone/db"
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func GetPhotos(c *gin.Context){
	user := utils.GetSession(c,"user")
	if user == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": "请登录！",
		})
	}else {
		photos,_ := QueryPhotosByUserName((user).(*model.User).UserName)
		c.HTML(http.StatusOK, "photo.html", gin.H{
			"msg": "Say",
			"user": user,
			"photos": photos,
		})
	}
}

//创建相册
func CreatePhoto(c *gin.Context)  {
	user := utils.GetSession(c,"user")
	u,_ := (user).(*model.User)
	photoName := c.DefaultPostForm("photoName","")
	createTime := utils.GetNowTimeShort()
	photo := model.Photo{
		Id:utils.GetUUID(),
		PhotoName:photoName,
		CreateTime:createTime,
		Creater:u.UserName,
	}
	if rdb := db.DB().Create(&photo); rdb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "创建失败!",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"mag": "创建成功!",
		})
	}
}

//查询相册
func QueryPhotosByUserName(name string) ([]model.Photo, error){
	photos := []model.Photo{}
	if rdb := db.DB().Where("creater = ?",name).Order("create_time desc").Find(&photos); rdb.Error != nil {
		return nil, rdb.Error
	}else {
		i := 0
		for _,photo := range photos{
			photos[i].Pictures,_ = QueryPicByPhotoId(photo.Id)
			i++
		}
		return photos,nil
	}
}

//根据相册id查询相册中照片
func QueryPicByPhotoId(photoId string) ([]model.Picture, error){
	pics := []model.Picture{}
	if rdb := db.DB().Where("photo_id = ?",photoId).Find(&pics); rdb.Error != nil {
		return nil, rdb.Error
	}else {
		return pics,nil
	}
}

//上传照片
func UploadPicture(c *gin.Context){
	user := utils.GetSession(c,"user")
	dirName := user.(*model.User).UserName
	path := "static/upload/"+dirName
	if !(utils.IsDirExist(path)){
		err := os.MkdirAll(path, os.ModePerm);
		if err != nil {
			fmt.Println(err)
		}
	}
	file, header, err := c.Request.FormFile("pic-upload")
	photoId := c.DefaultPostForm("photoId","")
	picName := header.Filename
	picSize := header.Size
	//存储的文件名加上时间戳避免相同名称被覆盖
	newName := utils.GetNowTimeStr()+picName
	picUrl := path+"/"+newName
	//创建文件
	out, err := os.Create(picUrl)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return
	}
	//测试一下
	fmt.Println("---------FILE---------")
	fmt.Println(file)
	fmt.Println("---------OUT---------")
	fmt.Println(out)
	b, err := ioutil.ReadFile(picUrl)
	fmt.Println("---------READ---------")
	fmt.Println(b)

	picture := model.Picture{
		Id:utils.GetUUID(),
		PicName:picName,
		PicSize:picSize,
		PicUrl:picUrl,
		PhotoId:photoId,
	}
	if rdb := db.DB().Create(&picture); rdb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "上传失败!",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"mag": "上传成功!",
		})
	}
}
