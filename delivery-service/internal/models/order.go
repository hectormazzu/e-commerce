package models

type Order struct {
	OrderID string      `json:"order_id"`
	Status  OrderStatus `json:"status"`
}
