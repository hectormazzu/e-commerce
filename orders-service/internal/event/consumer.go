package event

import (
	"encoding/json"
	"log"
	"orders-service/internal/db"
	"orders-service/internal/models"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func StartConsumer() {
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	log.Println("Connected to NATS")

	// Subscribe to the "delivery.events" subject
	nc.Subscribe("delivery.events", func(m *nats.Msg) {
		var DeliveryEvent models.DeliveryEvent
		if err := json.Unmarshal(m.Data, &DeliveryEvent); err != nil {
			log.Printf("Error parsing event: %v", err)
			return
		}

		log.Printf("[ORDERS-SERVICE] Received event for order %s with status %s\n", DeliveryEvent.OrderID, DeliveryEvent.Status)

		// Update the order status in the database
		err := updateOrderStatus(DeliveryEvent.OrderID, models.OrderStatus(DeliveryEvent.Status))
		if err != nil {
			log.Printf("Failed to update status for order %s: %v", DeliveryEvent.OrderID, err)
		}
	})

	select {} // Block forever
}

// Helper function to update the order status in the database
func updateOrderStatus(orderID string, status models.OrderStatus) error {
	var order models.Order
	if err := db.DB.First(&order, "id = ?", orderID).Error; err != nil {
		return err
	}

	order.Status = status
	if err := db.DB.Save(&order).Error; err != nil {
		return err
	}

	// Create a new OrderStatusHistory entry
	statusHistory := models.OrderStatusHistory{
		OrderID:   order.ID,
		Status:    status,
		Timestamp: time.Now(),
	}
	if err := db.DB.Create(&statusHistory).Error; err != nil {
		return err
	}

	log.Printf("Updated order %s to status %s\n", orderID, status)
	return nil
}
