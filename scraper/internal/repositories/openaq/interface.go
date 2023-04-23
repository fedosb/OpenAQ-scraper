package openaq

import "TPBDM/scraper/internal/entities"

type Repository interface {
	GetMeasurements(query QueryContract) ([]entities.Measurement, error)
	GetMeasurementsCount(query QueryContract) (int, error)
}
