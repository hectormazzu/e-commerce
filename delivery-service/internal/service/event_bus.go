package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn

// Initialize the NATS connection
func InitEventBus() error {
	var err error
	natsConnection, err = nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		return fmt.Errorf("failed to connect to NATS: %w", err)
	}
	log.Println("Connected to NATS")
	return nil
}

// PublishEvent publishes an event to the specified subject
func PublishEvent(subject string, event interface{}) error {
	if natsConnection == nil {
		return fmt.Errorf("NATS connection is not initialized")
	}

	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	err = natsConnection.Publish(subject, data)
	if err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	return nil
}
