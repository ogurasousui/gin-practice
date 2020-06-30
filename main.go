package main

import (
	"app/controller"
	"app/db"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db.Connection()
	defer db.DB().Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", controller.NewUser().List)
	r.GET("/users/:id", controller.NewUser().Get)
	r.POST("/users", controller.NewUser().Create)
	r.PUT("/users/:id", controller.NewUser().Update)
	r.DELETE("/users/:id", controller.NewUser().Delete)

	r.Run()
}
