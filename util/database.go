package util

import (
	"AnGO/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "ango"
)

// Creates a conection to DB, and return the session address
func GetConnection() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")
	db.AutoMigrate(&models.Province{})
	return db
}
