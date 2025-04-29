package event

import (
	common_models "cfotech/common/models"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func StartConsumer() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	log.Println("Connected to NATS")

	nc.Subscribe("delivery.events", func(m *nats.Msg) {
		var event common_models.DeliveryEvent
		if err := json.Unmarshal(m.Data, &event); err != nil {
			log.Printf("Error parsing event: %v", err)
			return
		}

		log.Printf("[NOTIFICATION] Order %s is now %s\n", event.OrderID, event.Status)

		// Send email notification
		err = sendEmailNotification(event)
		if err != nil {
			log.Printf("Failed to send email notification for order %s: %v", event.OrderID, err)
		}
	})

	select {} // block forever
}

// Helper function to send an email notification
func sendEmailNotification(event common_models.DeliveryEvent) error {
	// Replace with actual email sending logic
	log.Printf("Sending email notification for order %s with status %s\n", event.OrderID, event.Status)
	return nil
}
