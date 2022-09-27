package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	gorm.Model
	// ID       int    `json:"id"`
	Name   string `json:"name" form:"name"`
	Author string `json:"author" form:"author"`
	Page   int    `json:"page" form:"page"`
}
