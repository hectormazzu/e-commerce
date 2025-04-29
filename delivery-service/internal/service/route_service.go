package service

import (
	"delivery-service/internal/models"
	"delivery-service/internal/repository"
	"errors"
	"log"
	"sync"
)

func CreateRoute(vehicleID uint, driverID uint, Orders []models.Order) (*models.Route, error) {
	route := &models.Route{
		VehicleID: vehicleID,
		DriverID:  driverID,
		Orders:    Orders,
	}

	err := repository.CreateRoute(route)
	if err != nil {
		return nil, err
	}

	// Publish an event for route creation
	// event := models.DeliveryEvent{
	// 	OrderID: "", // No specific order for route creation
	// 	Status:  "RouteCreated",
	// 	Details: "Route created successfully",
	// }

	// if err := PublishEvent("delivery.events", event); err != nil {
	// 	log.Printf("Failed to publish route creation event: %v", err)
	// }

	return route, nil
}

func AddOrderToRoute(routeID, orderID string) (*models.Route, error) {
	route, err := repository.GetRoute(routeID)
	if err != nil {
		return nil, errors.New("route not found")
	}

	// Check if the order is already in the route
	for _, order := range route.Orders {
		if order.OrderID == orderID {
			return nil, errors.New("order is already in the route")
		}
	}

	route.Orders = append(route.Orders, models.Order{OrderID: orderID, Status: models.StatusPending})
	err = repository.UpdateRoute(route)
	if err != nil {
		return nil, err
	}

	// Publish an event for the added order
	event := models.DeliveryEvent{
		OrderID: orderID,
		Status:  string(models.StatusPending),
		Details: "Order added to route",
	}
	if err := PublishEvent("delivery.events", event); err != nil {
		log.Printf("Failed to publish order added event: %v", err)
	}

	return route, nil
}

func GetRoute(routeID string) (*models.Route, error) {
	// Fetch the route from the repository
	route, err := repository.GetRoute(routeID)
	if err != nil {
		return nil, errors.New("route not found")
	}

	// Concurrent fetching of order statuses
	// Use a WaitGroup to fetch order statuses in parallel
	var wg sync.WaitGroup
	mu := sync.Mutex{} // Mutex to protect shared data
	for i := range route.Orders {
		wg.Add(1)
		go func(order *models.Order) {
			defer wg.Done()

			// Call the orders-service to get the current status
			status, err := GetOrderStatus(order.OrderID)
			if err != nil {
				log.Printf("Failed to fetch status for order %s: %v", order.OrderID, err)
				status = "unknown" // Default status if the request fails
			}

			// Safely update the order's status
			mu.Lock()
			order.Status = models.OrderStatus(status)
			mu.Unlock()
		}(&route.Orders[i])
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Update the route in the repository
	err = repository.UpdateRoute(route)
	if err != nil {
		return nil, err
	}

	return route, nil
}

func StartRoute(routeID string) error {
	route, err := repository.GetRoute(routeID)
	if err != nil {
		return errors.New("route not found")
	}

	for i := range route.Orders {
		route.Orders[i].Status = models.StatusDispached

		// Publish an event for each dispatched order
		event := models.DeliveryEvent{
			OrderID: route.Orders[i].OrderID,
			Status:  string(models.StatusDispached),
			Details: "Order dispatched",
		}
		if err := PublishEvent("delivery.events", event); err != nil {
			log.Printf("Failed to publish order dispatched event for order %s: %v", route.Orders[i].OrderID, err)
		}
	}

	return repository.UpdateRoute(route)
}

func DeliverOrder(routeID string, orderID string) error {
	route, err := repository.GetRoute(routeID)
	if err != nil {
		return errors.New("route not found")
	}

	// Find the specific order in the route
	var order *models.Order
	log.Printf("Orders in route: %v\n", route.Orders)
	log.Printf("Looking for order %s in route %s\n", orderID, routeID)
	for i := range route.Orders {
		if route.Orders[i].OrderID == orderID {
			order = &route.Orders[i]
			break
		}
	}

	if order == nil {
		return errors.New("order not found in the route")
	}

	// Update the order status to StatusDelivered
	order.Status = models.StatusDelivered

	// Publish an event to NATS for the delivered order
	event := models.DeliveryEvent{
		OrderID: order.OrderID,
		Status:  string(models.StatusDelivered),
		Details: "Order delivered",
	}
	if err := PublishEvent("delivery.events", event); err != nil {
		log.Printf("Failed to publish order dispatched event for order %s: %v", order.OrderID, err)
	}

	return repository.UpdateRoute(route)
}

func UpdateRoute(route *models.Route) error {
	// Validate the route object
	if route == nil {
		return errors.New("route cannot be nil")
	}

	// Update the route in the repository
	err := repository.UpdateRoute(route)
	if err != nil {
		return err
	}

	return nil
}
