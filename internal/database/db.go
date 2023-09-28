package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "os"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        panic("A variável de ambiente DATABASE_URL não está definida.")
    }

    db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    DB = db
    return db, nil
}
