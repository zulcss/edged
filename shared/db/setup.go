package db

import (
	"fmt"
	"log"
	"gorm.io/gorm"
    	"gorm.io/driver/sqlite"
	
	"github.com/spf13/viper"

	"github.com/zulcss/edged/shared/model"
)

var DB *gorm.DB

func InitDatabase() {
	log.Println("Initialize database...")

	dbpath := viper.GetString("default.db_path")
	log.Printf("Using database: %s", dbpath)

	database, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s", dbpath)), &gorm.Config{})
    	if err != nil {
            panic("Failed to open database")
    	}

	// Create datbases
	database.AutoMigrate(&model.Site{})

	DB = database
}
