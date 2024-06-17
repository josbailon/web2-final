package config

import (
    "log"
    "os"
)

var (
    DatabaseDriver     string
    DatabaseConnection string
    JwtSecret          string
)

func LoadConfig() {
    DatabaseDriver = os.Getenv("DATABASE_DRIVER")
    if DatabaseDriver == "" {
        log.Fatal("DATABASE_DRIVER environment variable is required")
    }

    DatabaseConnection = os.Getenv("DATABASE_CONNECTION")
    if DatabaseConnection == "" {
        log.Fatal("DATABASE_CONNECTION environment variable is required")
    }

    JwtSecret = os.Getenv("JWT_SECRET")
    if JwtSecret == "" {
        log.Fatal("JWT_SECRET environment variable is required")
    }
}
