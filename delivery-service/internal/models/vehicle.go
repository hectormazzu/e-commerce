package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	// ID       string `json:"id" gorm:"primaryKey"`
	Make         string `json:"make"`
	VehicleModel string `json:"model"`
	Year         int    `json:"year"`
	Capacity     int    `json:"capacity"`
	// Route     Route  `json:"routes" gorm:"foreignKey:VehicleID"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}
