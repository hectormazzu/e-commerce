package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func GetOrderStatus(orderID string) (string, error) {
	// url := fmt.Sprintf("http://host.docker.internal:8081/orders/%s/status", orderID)
	url := fmt.Sprintf("%s/orders/%s/status", os.Getenv("ORDERS_SERVICE_URL"), orderID)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch order status: %s", resp.Status)
	}

	var response struct {
		OrderID string `json:"order_id"`
		Status  string `json:"current_status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	log.Printf("Fetched order status: %s\n", response)
	return response.Status, nil
}
