package repository

import (
	"delivery-service/internal/db"
	"delivery-service/internal/models"
)

func CreateRoute(route *models.Route) error {
	return db.DB.Preload("Vehicle").Preload("Driver").Create(route).First(route).Error
}

func GetRoute(routeID string) (*models.Route, error) {
	var route models.Route
	if err := db.DB.Preload("Vehicle").Preload("Driver").First(&route, "id = ?", routeID).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func UpdateRoute(route *models.Route) error {
	return db.DB.Save(route).Error
}
