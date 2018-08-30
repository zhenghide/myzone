package main

import (
	"myzone/router"
	"github.com/spf13/viper"
	"flag"
	"strings"
	"myzone/db"
	"encoding/gob"
	"myzone/model"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

var (
	fConfig = flag.String("config", "config.yaml", "configuration file to load")
)

func main() {
	//设置配置
	viper.AddConfigPath("config")
	*fConfig = strings.Replace(*fConfig, ".yaml", "", 1)
	viper.SetConfigName(*fConfig)

	err := viper.ReadInConfig()
	if err != nil {
		panic("can not read config: " + err.Error())
	}

	err = db.Init()
	if err != nil {
		panic("can not connect to db: " + err.Error())
	}

	//初始化路由
	router := router.InitRouter()

	//读取端口配置
	port := ":" + viper.GetString("port.port")

	router.Run(port)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		db.CloseDB()
		fmt.Println("myzone exit.....")
		os.Exit(0)
	}()
}

func init()  {
	//使用session需要注册结构体
	gob.Register(&model.User{})
}
