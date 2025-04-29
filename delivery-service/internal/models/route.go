package models

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	// ID        string    `json:"id"`
	VehicleID uint    `json:"vehicle_id"`
	Vehicle   Vehicle `json:"vehicle" gorm:"foreignKey:VehicleID"`
	DriverID  uint    `json:"driver_id"`
	Driver    Driver  `json:"driver" gorm:"foreignKey:DriverID"`
	Orders    []Order `json:"orders" gorm:"serializer:json"` // Store as JSON in SQLite
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}
