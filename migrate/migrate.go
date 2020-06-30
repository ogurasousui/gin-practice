package main

import (
	"app/db"
	"app/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db.Connection()
	defer db.DB().Close()

	db.DB().AutoMigrate(&model.User{})
	db.DB().AutoMigrate(&model.UserName{})
}
