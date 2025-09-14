package repository

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=testdb port=5432 sslmode=disable"
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ { // пробуем 10 раз
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			return db, nil
		}
		fmt.Println("DB not ready, retrying in 2s...")
		time.Sleep(2 * time.Second)
	}

	return nil, err
}
