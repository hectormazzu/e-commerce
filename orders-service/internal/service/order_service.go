package service

import (
	"errors"
	"orders-service/internal/models"
	"orders-service/internal/repository"
)

func GetOrderStatus(orderID string) (*models.Order, error) {
	// Fetch the order from the repository
	order, err := repository.GetOrder(orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	// Check if the order has a status history
	if len(order.History) == 0 {
		return nil, errors.New("no status history found for this order")
	}

	return order, nil
}
