package main

import (
	"first_go_project/internal/events/orders"
	handlers2 "first_go_project/internal/handlers"
	models2 "first_go_project/internal/models"
	"first_go_project/internal/redis"
	"first_go_project/internal/repository"
	"first_go_project/internal/routes"
	"first_go_project/internal/seed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	db, err := repository.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models2.User{}, &models2.Product{}, &models2.Order{}, &models2.OrderItem{})

	if err := seed.Run(db); err != nil {
		log.Fatal("Failed seeding data: ", err)
	}

	redis.InitRedis()

	// слушаем события
	orders.ListenOrderCreated()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	handlers2.RegisterAuthRoutes(r)
	routes.UserRoutes(r)
	handlers2.RegisterOrderRoutes(r)

	r.Run(":8080")
}
