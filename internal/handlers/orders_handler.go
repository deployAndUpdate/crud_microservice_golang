// internal/handlers/orders_handler.go
package handlers

import (
	orders "first_go_project/internal/events/orders"
	"first_go_project/internal/middleware"
	models "first_go_project/internal/models"
	repository "first_go_project/internal/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine) {
	r.GET("/orders", middleware.JWTMiddleware(), getOrders)
	r.POST("/orders", middleware.JWTMiddleware(), createOrder)
	r.GET("/orders/:id", middleware.JWTMiddleware(), getOrderById)
	r.GET("/users/:id/orders", middleware.JWTMiddleware(), getOrdersByUserId)
	r.PUT("/orders/:id", middleware.JWTMiddleware(), updateOrder)
	r.DELETE("/orders/:id", middleware.JWTMiddleware(), deleteOrder)
}

func createOrder(c *gin.Context) {
	var payload struct {
		Name  string `json:"name"`
		Items []struct {
			ProductID uint `json:"product_id"`
			Quantity  uint `json:"quantity"`
		} `json:"items"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := uid.(uint)

	order := &models.Order{
		Name:   payload.Name,
		UserID: &userID,
	}

	var items []models.OrderItem
	for _, it := range payload.Items {
		items = append(items, models.OrderItem{
			ProductID: it.ProductID,
			Quantity:  it.Quantity,
		})
	}

	if err := repository.CreateOrderWithItems(order, items); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Публикуем событие
	orders.PublishOrderCreated(orders.OrderCreatedEvent{
		OrderID: order.ID,
		UserID:  userID,
		Status:  "NEW",
	})

	c.JSON(http.StatusCreated, order)
}

func getOrders(c *gin.Context) {
	ordersRes, err := repository.GetOrders()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ordersRes)
}

func getOrderById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := repository.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func getOrdersByUserId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := repository.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ordersRes, err := repository.GetOrdersByUserID(uint(id))
	if err != nil || len(ordersRes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Orders of %s not found", user.Name),
		})
		return
	}

	c.JSON(http.StatusOK, ordersRes)
}

func updateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := repository.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.UpdateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func deleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.DeleteOrder(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
