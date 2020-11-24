package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {

	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/wsadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(new(Admin))

	admin := Admin{
		Username: "瞄瞄🌞",
	}

	db.Save(&admin)
}

func DB() *gorm.DB {
	return db
}
