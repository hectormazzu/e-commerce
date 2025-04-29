package handlers

import (
	"log"
	"net/http"
	"orders-service/internal/service"

	"github.com/gin-gonic/gin"
)

func GetOrderStatus(c *gin.Context) {
	id := c.Param("id")

	// Call the service layer to get the order status
	order, err := service.GetOrderStatus(id)
	if err != nil {
		log.Printf("Error fetching order status: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the order status and history
	c.JSON(http.StatusOK, gin.H{
		"order_id":       order.ID,
		"current_status": order.Status,
		"history":        order.History,
	})
}
