package test

import (
	"net/http"
	"net/http/httptest"
	"orders-service/internal/db"
	"orders-service/internal/handlers"

	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetOrderStatus(t *testing.T) {
	db.Init()
	r := gin.Default()
	r.GET("/orders/:id/status", handlers.GetOrderStatus)

	req, _ := http.NewRequest("GET", "/orders/order1/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", w.Code)
	}
}
