package main

import (
	"delivery-service/internal/db"
	"delivery-service/internal/handlers"
	"delivery-service/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	// if err := db.Init(os.Getenv("DB_URL")); err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.CloseDB()
	db.Init()

	// Initialize the event bus
	if err := service.InitEventBus(); err != nil {
		log.Fatalf("Failed to initialize event bus: %v", err)
	}

	r := gin.Default()
	// Register routes
	r.POST("/routes", handlers.CreateRoute)
	r.GET("/routes/:id", handlers.GetRoute)
	r.POST("/routes/:id/orders", handlers.AddOrderToRoute)
	r.POST("/routes/:id/start", handlers.StartRoute)
	r.POST("/routes/:id/orders/:order_id/deliver", handlers.DeliverOrder)

	log.Println("Delivery service running on port:", os.Getenv("DELIVERY_SERVICE_PORT"))
	r.Run(":" + os.Getenv("DELIVERY_SERVICE_PORT"))
	// r.Run(":8082") // listen and serve on
}
