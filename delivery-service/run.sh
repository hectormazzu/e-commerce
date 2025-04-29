#!/bin/bash
ORDERS_SERVICE_URL=http://localhost:8081 NATS_URL=nats://localhost:4222 DELIVERY_SERVICE_PORT=8082 go run cmd/main.go