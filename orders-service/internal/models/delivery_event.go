package models

type DeliveryEvent struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
	Details string `json:"details"`
}
