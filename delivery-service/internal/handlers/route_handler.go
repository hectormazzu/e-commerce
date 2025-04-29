package handlers

import (
	"delivery-service/internal/models"
	"delivery-service/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoute(c *gin.Context) {
	var input struct {
		Vehicle uint           `json:"vehicle_id" binding:"required"`
		Driver  uint           `json:"driver_id"  binding:"required"`
		Orders  []models.Order `json:"orders"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	route, err := service.CreateRoute(input.Vehicle, input.Driver, input.Orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create route"})
		return
	}

	c.JSON(http.StatusCreated, route)
}

func GetRoute(c *gin.Context) {
	routeID := c.Param("id")
	route, err := service.GetRoute(routeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
		return
	}

	c.JSON(http.StatusOK, route)
}

func AddOrderToRoute(c *gin.Context) {
	routeID := c.Param("id")
	var input struct {
		OrderID string `json:"order_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	route, err := service.AddOrderToRoute(routeID, input.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, route)
}

func StartRoute(c *gin.Context) {
	routeID := c.Param("id")
	err := service.StartRoute(routeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Route started successfully"})
}

func DeliverOrder(c *gin.Context) {
	routeID := c.Param("id")
	orderID := c.Param("order_id")
	log.Printf("Delivering order %s for route %s\n", orderID, routeID)
	err := service.DeliverOrder(routeID, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order delivered successfully"})
}
