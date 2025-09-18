package models

import (
    "log"
    "os"
    "time"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
    // Build connection string from env variables
    dsn := "host=localhost user=postgres dbname=hustle port=5432 sslmode=disable TimeZone=UTC"
    
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    sqlDB, err := database.DB()
    if err != nil {
        panic("Failed to get underlying sql.DB")
    }
    
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

	// Make sure to include all models to migrate here
    err = database.AutoMigrate(&User{})
    if err != nil {
        log.Printf("Failed to migrate database: %v", err)
        panic("Database migration failed!")
    }

    DB = database
    log.Println("Database connected successfully")
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}