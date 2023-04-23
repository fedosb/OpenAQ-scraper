package openaq

import "TPBDM/scraper/internal/entities"

type Repository interface {
	GetMeasurements() ([]entities.Measurement, error)
	GetMeasurementsCount() (int, error)
}
