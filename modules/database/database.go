package database

import (
	"fmt"
	"github.com/arifbugaresa/go-hexa/config"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(config *config.AppConfig) *gorm.DB {
	var dsn string

	log.Info("Starting on " + config.AppEnvironment)
	if config.AppEnvironment == "development" {
		dsn = fmt.Sprintf(`host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta search_path=%s`,
			config.DBHost, config.DBUsername, config.DBPassword, config.DBPort, config.DBName, config.SchemaName)
	} else if config.AppEnvironment == "sandbox" {
		dsn = fmt.Sprintf(`host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Jakarta`,
			config.DBHost, config.DBUsername, config.DBPassword, config.DBPort, config.DBName)
	}

	dbGormPostgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Failed")
		fmt.Println("DB Connection Failed")
	} else {
		log.Info("DB Connection Success")
		fmt.Println("DB Connection Success")
	}

	return dbGormPostgres
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()

	log.Info("DB Connection Close")
	fmt.Println("DB Connection Close")
}
