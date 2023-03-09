package db

import (
	"log"
	"gorm.io/gorm"
    "gorm.io/driver/sqlite"

	"github.com/zulcss/edged/shared/model"
)

var DB *gorm.DB

func InitDatabase() {
	log.Println("Initialize database...")

	database, err := gorm.Open(sqlite.Open("edged.db"), &gorm.Config{})
    	if err != nil {
            panic("Failed to open database")
    	}

	// Create datbases
	database.AutoMigrate(&model.Site{})

	DB = database
}
