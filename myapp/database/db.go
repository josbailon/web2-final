package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "log"
    "parking-management/config"
    "parking-management/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    if config.DatabaseDriver == "sqlite" {
        DB, err = gorm.Open(sqlite.Open(config.DatabaseConnection), &gorm.Config{})
    } else {
        log.Fatal("Unsupported database driver")
    }
    
    if err != nil {
        log.Fatal("Failed to connect to database")
    }

    // Migrate the schema
    DB.AutoMigrate(&models.Vehicle{}, &models.Payment{})
}
