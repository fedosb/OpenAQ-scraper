package measurements

import "TPBDM/scraper/internal/entities"

type Repository interface {
	CreateMeasurement(measurement entities.Measurement) error
	GetMeasurementsList() ([]entities.Measurement, error)
}
