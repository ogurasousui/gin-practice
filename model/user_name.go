package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserName struct {
	gorm.Model
	UserID uint
	Name   string
}
