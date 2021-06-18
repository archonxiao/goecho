package dao

import (
	"goecho/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect : Database connect
func Connect() *gorm.DB {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/goecho?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Article{})
	return db
}

func GetInstance() *gorm.DB {
	return db
}
