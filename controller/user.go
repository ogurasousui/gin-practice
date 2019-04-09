package controller

import (
	"app/db"
	"app/model"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var user model.User
	result := db.First(&user, c.Param("id")).Related(&user.UserName)
	c.JSON(200, result.Value)
}

func List(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var users []model.User
	result := db.Preload("UserName").Find(&users)
	c.JSON(200, result.Value)
}

func Create(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var user = model.User{}
	db.Create(&user)

	var userName = model.UserName{}
	c.BindJSON(&userName)

	userName.UserID = user.ID
	db.Create(&userName)
}
