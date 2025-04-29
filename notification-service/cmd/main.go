package main

import (
	"log"
	"notification-service/internal/event"
)

func main() {
	log.Println("Starting Notification Service...")
	event.StartConsumer()
}
