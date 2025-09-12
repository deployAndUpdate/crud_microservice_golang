package repository

import (
	"first_go_project/internal/models"
)

func CreateOrder(order *models.Order) error {
	return DB.Create(order).Error
}

func GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := DB.Preload("User").Find(&orders).Error // Preload подтянет данные юзера
	return orders, err
}

func GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := DB.Preload("User").First(&order, id).Error
	return &order, err
}

func GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func UpdateOrder(order *models.Order) error {
	return DB.Save(order).Error
}

func DeleteOrder(id uint) error {
	return DB.Delete(&models.Order{}, id).Error
}
