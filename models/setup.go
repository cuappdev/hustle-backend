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
    host := getEnv("DB_HOST", "localhost")
    user := getEnv("DB_USER", "postgres")
    dbname := getEnv("DB_NAME", "hustle")
    port := getEnv("DB_PORT", "5432")
    password := getEnv("DB_PASSWORD", "")
    sslmode := getEnv("DB_SSLMODE", "require")
    
    dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode + " TimeZone=UTC"
    
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    sqlDB, err := database.DB()
    if err != nil {
        return err
    }
    
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

	// Make sure to include all models to migrate here
    err = database.AutoMigrate(&User{})
    if err != nil {
        log.Printf("Failed to migrate database: %v", err)
        return err
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