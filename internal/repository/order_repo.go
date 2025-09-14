package repository

import (
	models2 "first_go_project/internal/models"

	"gorm.io/gorm"
)

func CreateOrder(order *models2.Order) error {
	return DB.Create(order).Error
}

func GetOrders() ([]models2.Order, error) {
	var orders []models2.Order
	err := DB.Preload("User").Find(&orders).Error // Preload подтянет данные юзера
	return orders, err
}

func GetOrderByID(id uint) (*models2.Order, error) {
	var order models2.Order
	err := DB.Preload("User").First(&order, id).Error
	return &order, err
}

func GetOrdersByUserID(userID uint) ([]models2.Order, error) {
	var orders []models2.Order
	err := DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func UpdateOrder(order *models2.Order) error {
	return DB.Save(order).Error
}

func DeleteOrder(id uint) error {
	return DB.Delete(&models2.Order{}, id).Error
}

func CreateOrderWithItems(order *models2.Order, items []models2.OrderItem) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// проверка продуктов + создание позиций
		for i := range items {
			// убедиться, что продукт существует
			var prod models2.Product
			if err := tx.First(&prod, items[i].ProductID).Error; err != nil {
				return err // rollback
			}
			items[i].OrderID = order.ID
			if err := tx.Create(&items[i]).Error; err != nil {
				return err
			}
		}

		// посчитать и обновить сумму
		var total uint = 0
		for _, it := range items {
			var p models2.Product
			if err := tx.First(&p, it.ProductID).Error; err != nil {
				return err
			}
			total += p.Price * it.Quantity
		}
		order.Amount = total
		if err := tx.Save(order).Error; err != nil {
			return err
		}

		return nil
	})
}
