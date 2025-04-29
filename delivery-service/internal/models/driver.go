package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	// ID        string `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	// CreatedAt string `json:"created_at"`
	// UpdatedAt string `json:"updated_at"`
	// Route     Route  `json:"routes" gorm:"foreignKey:DriverID"`
	// Orders    []Order `json:"orders" gorm:"foreignKey:DriverID"`
	// Vehicles  []Vehicle `json:"vehicles" gorm:"foreignKey:DriverID"`

}
