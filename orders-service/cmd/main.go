package main

import (
	"log"
	"orders-service/internal/db"
	"orders-service/internal/event"
	"orders-service/internal/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.Init()

	// Start the NATS consumer
	go event.StartConsumer()

	r := gin.Default()
	r.GET("/orders/:id/status", handlers.GetOrderStatus)

	r.Run(":" + os.Getenv("ORDERS_SERVICE_PORT")) // listen and serve on port 8081
	// r.Run(":8081") // listen and serve on port 8081
	log.Println("Orders service is running...")
	select {} // Block forever
}
