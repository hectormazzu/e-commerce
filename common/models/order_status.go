package models

type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"    // Order is added to route
	StatusDispached OrderStatus = "DISPATCHED" // Order is dispatched to the delivery person
	StatusDelivered OrderStatus = "DELIVERED"  // Order is delivered to the customer
	StatusCanceled  OrderStatus = "CANCELED"   // Order is canceled by the customer or system
)
