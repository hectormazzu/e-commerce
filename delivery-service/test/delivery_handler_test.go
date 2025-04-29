package test

import (
	"bytes"
	"delivery-service/internal/db"
	"delivery-service/internal/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	db.Init() // Initialize the in-memory database
	r := gin.Default()
	r.POST("/routes", handlers.CreateRoute)
	r.POST("/routes/:id/orders", handlers.AddOrderToRoute)
	return r
}

func TestCreateRoute(t *testing.T) {
	r := setupRouter()

	body := `{"vehicle": "Truck 1", "driver": "Alice"}`
	req, _ := http.NewRequest("POST", "/routes", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"vehicle":"Truck 1"`)
	assert.Contains(t, w.Body.String(), `"driver":"Alice"`)
}

func TestAddOrderToRoute(t *testing.T) {
	r := setupRouter()

	// Create a route first
	routeBody := `{"vehicle": "Truck 1", "driver": "Alice"}`
	req, _ := http.NewRequest("POST", "/routes", bytes.NewBuffer([]byte(routeBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Extract the route ID from the response
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	routeID := response["id"].(string)

	// Add an order to the route
	orderBody := `{"order_id": "order123"}`
	req, _ = http.NewRequest("POST", "/routes/"+routeID+"/orders", bytes.NewBuffer([]byte(orderBody)))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"order123"`)
}
