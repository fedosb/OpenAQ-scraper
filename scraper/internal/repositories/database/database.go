package database

import (
	"TPBDM/scraper/internal/repositories/database/measurements"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Container struct {
	AirQualityRecords measurements.Repository
}

// New ...
func New() (*Container, error) {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		"postgres",
		"securepasswd",
		"tpbdm",
		"localhost",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&measurements.MeasurementModel{})
	if err != nil {
		return nil, err
	}

	return &Container{
		AirQualityRecords: measurements.NewMeasurementsRepository(db),
	}, nil
}
