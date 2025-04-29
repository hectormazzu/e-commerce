package db

import (
	"log"
	"orders-service/internal/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB.AutoMigrate(&models.Order{}, &models.OrderStatusHistory{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Order{}, &models.OrderStatusHistory{})

	// avoid error no such table: orders
	// Set the maximum number of open connections to 1 to avoid connection pool exhaustion
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB from gorm.DB: ", err)
	}
	sqlDB.SetMaxOpenConns(1)

	seedOrders()
}

func seedOrders() {
	orders := []models.Order{
		{ID: "order1", CreatedAt: time.Now().Add(-24 * time.Hour), Status: "PENDING"},
		{ID: "order2", CreatedAt: time.Now().Add(-24 * time.Hour), Status: "PENDING"},
		{ID: "order3", CreatedAt: time.Now().Add(-48 * time.Hour), Status: "PENDING"},
	}

	for _, order := range orders {
		DB.Create(&order)
		status := models.OrderStatusHistory{
			OrderID:   order.ID,
			Status:    order.Status,
			Timestamp: time.Now(),
		}
		DB.Create(&status)
	}
}
