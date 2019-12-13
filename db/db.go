package db

import (
	"fmt"
	"github.com/mblode/gospel/app/models"
	"github.com/mblode/gospel/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func initialMigration(db *gorm.DB) {
	db.Debug().AutoMigrate(&models.User{})
}

func Init() {
	configuration := config.GetConfig()
	connect := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DbUsername, configuration.DbPassword, configuration.DbName)
	db, err = gorm.Open("mysql", connect)
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}

	initialMigration(db)
}

func GetDb() *gorm.DB {
	return db
}
