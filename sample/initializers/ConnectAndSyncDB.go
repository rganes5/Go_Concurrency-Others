package initializers

import (
	"fmt"
	"log"
	"sample/config"
	"sample/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func ConnectAndSyncDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.Host, config.Cfg.Port, config.Cfg.User, config.Cfg.Password, config.Cfg.DbName)

	var err error

	gormDB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate tables here
	// Example:
	err = gormDB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetGormDB() *gorm.DB {
	return gormDB
}
