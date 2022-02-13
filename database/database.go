package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB(){

	str := "root:1234@tcp(db:3306)/go_api?charset=utf8mb4&parseTime=True&loc=Local";

	database, err := gorm.Open(mysql.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
		
	}

	db = database

	config, _  := db.DB();

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

}

func GetDatabase() *gorm.DB {
	return db
}