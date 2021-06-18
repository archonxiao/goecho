package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Content datatypes.JSON
	Read int `gorm:"type:tinyint(1);default:0"`
}
