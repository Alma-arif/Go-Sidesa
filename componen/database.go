package componen

import (
	"fmt"
	"go-sidesa/app/users"
	"go-sidesa/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(conf config.Config) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DB.User, conf.DB.Pass, conf.DB.Host, conf.DB.Port, conf.DB.Name)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db

	db.AutoMigrate(&users.User{})

}
