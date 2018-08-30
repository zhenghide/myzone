package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

func Init() (err error) {
	db, err = gorm.Open("mysql", viper.GetString("db.server"))
	if err != nil {
		return
	}
	db.SingularTable(true)

	if viper.GetString("log.level") == "debug" {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	return
}

func DB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}