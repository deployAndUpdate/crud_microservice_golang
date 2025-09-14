package repository

import (
	"first_go_project/internal/models"
)

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

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
