package repository

import (
	"first_go_project/internal/models"
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
			DB = db // <--- присваиваем глобальной переменной
			return db, nil
		}
		fmt.Println("DB not ready, retrying in 2s...")
		time.Sleep(2 * time.Second)
	}

	return nil, err
}

func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := DB.Find(&users).Error
	return users, err
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := DB.First(&user, id).Error
	return &user, err
}

func UpdateUser(user *models.User) error {
	return DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return DB.Delete(&models.User{}, id).Error
}
