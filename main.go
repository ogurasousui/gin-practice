package main

import (
	"app/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", controller.List)
	r.GET("/users/:id", controller.Get)
	r.POST("/users", controller.Create)ｓｙ

	r.Run()
}
