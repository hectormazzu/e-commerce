package models

import (
	common_models "cfotech/common/models"
	"time"
)

type Order struct {
	ID        string                    `json:"id"`
	CreatedAt time.Time                 `json:"created_at"`
	Status    common_models.OrderStatus `json:"status"`
	History   []OrderStatusHistory      `json:"history"`
}

type OrderStatusHistory struct {
	ID        uint                      `json:"id"`
	OrderID   string                    `json:"order_id"`
	Status    common_models.OrderStatus `json:"status"`
	Timestamp time.Time                 `json:"timestamp"`
}
