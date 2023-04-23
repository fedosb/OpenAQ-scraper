package database

import (
	"TPBDM/scraper/config"
	"TPBDM/scraper/internal/repositories/database/measurements"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Container struct {
	Measurements measurements.Repository
}

// New ...
func New(config config.DBConfig) (*Container, error) {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		config.Username,
		config.Password,
		config.Database,
		config.Host,
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
		Measurements: measurements.NewMeasurementsRepository(db),
	}, nil
}
