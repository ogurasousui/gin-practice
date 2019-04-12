package controller

import (
	"app/db"
	"app/model"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (t *User) Get(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var user model.User
	result := db.First(&user, c.Param("id")).Related(&user.UserName)
	c.JSON(200, result.Value)
}

func (t *User) List(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var users []model.User
	result := db.Preload("UserName").Find(&users)
	c.JSON(200, result.Value)
}

func (t *User) Create(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var user model.User
	db.Create(&user)

	var userName model.UserName
	c.BindJSON(&userName)
	userName.UserID = user.ID
	db.Create(&userName)
}

func (t *User) Update(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var user model.User
	db.First(&user, c.Param("id")).Related(&user.UserName)
	if user.UserName.ID > 0 {
		c.BindJSON(&user.UserName)
		db.Save(&user.UserName)
	}
}

func (t *User) Delete(c *gin.Context) {
	db := db.Connection()
	defer db.Close()

	var user model.User
	db.First(&user, c.Param("id")).Related(&user.UserName)
	if user.ID > 0 {
		db.Delete(&user)
	}
	if user.UserName.ID > 0 {
		db.Delete(&user.UserName)
	}

	// こっちの消し方でも良い
	// db.Where("id = ?", c.Param("id")).Delete(&model.User{})
	// db.Where("user_id = ?", c.Param("id")).Delete(&model.UserName{})
}
