package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func Connection() *gorm.DB {
	var err error
	database, err = gorm.Open("mysql", "root:@tcp(db:3306)/gin_app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	database.LogMode(true)
	return database
}

func DB() *gorm.DB {
	return database
}
