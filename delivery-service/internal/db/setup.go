package db

import (
	"delivery-service/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto-migrate the Route model
	if err := DB.AutoMigrate(&models.Route{}, &models.Vehicle{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	seedVehicle()
	seedDriver()
}

func seedVehicle() {
	vehicle := models.Vehicle{Make: "Toyota", VehicleModel: "Corolla", Year: 2020, Capacity: 5}
	DB.Create(&vehicle)
}

func seedDriver() {
	vehicle := models.Driver{Name: "John Doe", Phone: "1234567890", Email: "JohnDoe@example.com"}
	DB.Create(&vehicle)
}
