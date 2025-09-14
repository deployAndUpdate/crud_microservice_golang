package seed

import (
	models2 "first_go_project/internal/models"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	rand.Seed(time.Now().UnixNano())

	// users
	for i := 0; i < 20; i++ {
		u := models2.User{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Role:  "user",
		}
		db.FirstOrCreate(&u, models2.User{Email: u.Email})
	}

	// products
	products := []models2.Product{
		{Name: "T-shirt", Price: 100},
		{Name: "Sneakers", Price: 500},
		{Name: "Hat", Price: 50},
		{Name: "Book", Price: 150},
	}
	for _, p := range products {
		db.FirstOrCreate(&p, models2.Product{Name: p.Name})
	}

	// orders & items
	var users []models2.User
	db.Find(&users)
	var prods []models2.Product
	db.Find(&prods)

	for i := 0; i < 20; i++ {
		user := users[i%len(users)]
		numItems := rand.Intn(3) + 1
		order := models2.Order{
			Name:   fmt.Sprintf("Order-%d", i),
			UserID: &user.ID,
			Amount: 0,
		}
		db.Create(&order)
		var total uint = 0
		for j := 0; j < numItems; j++ {
			p := prods[rand.Intn(len(prods))]
			qty := uint(rand.Intn(3) + 1)
			oi := models2.OrderItem{
				OrderID:   order.ID,
				ProductID: p.ID,
				Quantity:  qty,
			}
			db.Create(&oi)
			total += p.Price * qty
		}
		order.Amount = total
		db.Save(&order)
	}

	return nil
}
