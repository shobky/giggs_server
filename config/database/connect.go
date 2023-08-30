package database

import (
	"fmt"
	"log"

	"github.com/shobky/giggs/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s sslmode=disable",
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbUser,
		config.Config.DbPassword,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	fmt.Println("🚀 Connected Successfully to the Database")
}
