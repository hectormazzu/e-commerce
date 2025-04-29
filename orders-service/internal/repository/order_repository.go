package repository

import (
	"orders-service/internal/db"
	"orders-service/internal/models"
)

func GetOrder(orderID string) (*models.Order, error) {
	var order models.Order

	// Fetch the order from the database
	if err := db.DB.Preload("History").First(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
