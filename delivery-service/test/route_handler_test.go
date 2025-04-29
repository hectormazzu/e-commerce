package test

import (
	"bytes"
	common_models "cfotech/common/models"
	"delivery-service/internal/db"
	"delivery-service/internal/handlers"
	"delivery-service/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	db.Init() // Initialize the in-memory database
	r := gin.Default()
	// r.POST("/routes", handlers.CreateRoute)
	// r.POST("/routes/:id/orders/:order_id/deliver", handlers.DeliverOrder)
	r.POST("/routes", handlers.CreateRoute)
	r.GET("/routes/:id", handlers.GetRoute)
	r.POST("/routes/:id/orders", handlers.AddOrderToRoute)
	r.POST("/routes/:id/start", handlers.StartRoute)
	r.POST("/routes/:id/orders/:order_id/deliver", handlers.DeliverOrder)
	return r
}
func TestCreateRoute(t *testing.T) {
	r := setupRouter()

	// Test case: Successfully create a route
	reqBody := []byte(`{"vehicle_id":1,"driver_id":1,"orders":[]}`)
	req, _ := http.NewRequest("POST", "/routes", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"vehicle_id":1`)
	assert.Contains(t, w.Body.String(), `"driver_id":1`)

	// Test case: Missing required fields
	reqBody = []byte(`{"vehicle_id":1}`)
	req, _ = http.NewRequest("POST", "/routes", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error"`)
}

func TestGetRoute(t *testing.T) {
	r := setupRouter()

	// Seed a route
	route := models.Route{
		VehicleID: 1,
		DriverID:  1,
		Orders:    []models.Order{},
	}
	db.DB.Create(&route)

	// Test case: Successfully fetch a route
	req, _ := http.NewRequest("GET", "/routes/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"vehicle_id":1`)
	assert.Contains(t, w.Body.String(), `"driver_id":1`)

	// Test case: Route not found
	req, _ = http.NewRequest("GET", "/routes/999", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"Route not found"`)
}

func TestAddOrderToRoute(t *testing.T) {
	r := setupRouter()

	// Seed a route
	route := models.Route{
		VehicleID: 1,
		DriverID:  1,
		Orders:    []models.Order{},
	}
	db.DB.Create(&route)

	// Test case: Successfully add an order to a route
	reqBody := []byte(`{"order_id":"order123"}`)
	req, _ := http.NewRequest("POST", "/routes/1/orders", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"order_id":"order123"`)

	// Test case: Route not found
	req, _ = http.NewRequest("POST", "/routes/999/orders", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"route not found"`)
}

func TestStartRoute(t *testing.T) {
	r := setupRouter()

	// Seed a route with orders
	route := models.Route{
		VehicleID: 1,
		DriverID:  1,
		Orders: []models.Order{
			{OrderID: "order123", Status: "PENDING"},
		},
	}
	db.DB.Create(&route)

	// Test case: Successfully start a route
	req, _ := http.NewRequest("POST", "/routes/1/start", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"message":"Route started successfully"`)

	// Verify the order status is updated to DISPATCHED
	var updatedRoute models.Route
	db.DB.Preload("Orders").First(&updatedRoute, "id = ?", route.ID)
	assert.Equal(t, common_models.StatusDispached, updatedRoute.Orders[0].Status)

	// Test case: Route not found
	req, _ = http.NewRequest("POST", "/routes/999/start", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"route not found"`)
}
func TestDeliverOrder1(t *testing.T) {
	r := setupRouter()

	// Seed a route with an order
	route := models.Route{
		VehicleID: 1,
		DriverID:  1,
		Orders: []models.Order{
			{OrderID: "order123", Status: "DISPATCHED"},
		},
	}
	db.DB.Create(&route)

	// Test case: Successfully deliver an order
	req, _ := http.NewRequest("POST", "/routes/1/orders/order123/deliver", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"message":"Order delivered successfully"`)

	// Verify the order status is updated to DELIVERED
	var updatedRoute models.Route
	db.DB.Preload("Orders").First(&updatedRoute, "id = ?", route.ID)
	assert.Equal(t, common_models.StatusDelivered, updatedRoute.Orders[0].Status)

	// Test case: Order not found in the route
	req, _ = http.NewRequest("POST", "/routes/1/orders/order999/deliver", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"order not found in the route"`)
}
