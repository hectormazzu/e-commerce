package models

import (
	"time"
)

type Order struct {
	ID        string               `json:"id"`
	CreatedAt time.Time            `json:"created_at"`
	Status    OrderStatus          `json:"status"`
	History   []OrderStatusHistory `json:"history"`
}

type OrderStatusHistory struct {
	ID        uint        `json:"id"`
	OrderID   string      `json:"order_id"`
	Status    OrderStatus `json:"status"`
	Timestamp time.Time   `json:"timestamp"`
}
