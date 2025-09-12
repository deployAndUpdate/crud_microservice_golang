package main

import (
	"first_go_project/internal/models"
	"first_go_project/internal/repository"
	"first_go_project/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{}, &models.Order{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.UserRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
