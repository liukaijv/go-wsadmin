package models

import (
	"go-wsadmin/config"
	"go-wsadmin/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func InitDb() error {

	if db != nil {
		return nil
	}

	var err error
	dsn := config.Conf.DbDsn
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if config.Conf.DebugMode {
		db = db.Debug()
	}

	initModels()

	initSeeds()

	return nil
}

func initModels() error {
	return db.AutoMigrate(new(Admin))
}

func initSeeds() {

	if config.Conf.SuperAdmin != "" && config.Conf.SuperAdminPassword != "" {
		var (
			username = config.Conf.SuperAdmin
			password = utils.GeneratePassword(config.Conf.SuperAdminPassword)
		)
		var admin Admin
		err := db.Where("username = ?", username).First(&admin).Error
		if err == gorm.ErrRecordNotFound {
			db.Save(&Admin{Username: username, Password: password})
		} else {
			log.Print("管理员已初始化")
		}
	}

}
