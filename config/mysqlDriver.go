package config

import (
	_mUsers "altaproject/features/users/data"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_Username")
	dbPassword := os.Getenv("DB_Password")
	dbPort := os.Getenv("DB_Port")
	dbHost := os.Getenv("DB_Host")
	dbName := os.Getenv("DB_Name")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	// var e error
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
}
