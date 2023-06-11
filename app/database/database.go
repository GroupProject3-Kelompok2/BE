package database

import (
	"fmt"

	"github.com/GroupProject3-Kelompok2/BE/app/config"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var log = middlewares.Log()

func InitDatabase(config config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUSER, config.DBPASSWORD, config.DBHOST, config.DBPORT, config.DBNAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	log.Info("success connected to database")

	return db
}
