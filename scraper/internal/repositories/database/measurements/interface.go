package measurements

import "TPBDM/scraper/internal/entities"

type Repository interface {
	CreateMeasurement(measurement entities.Measurement) error
	GetMeasurementsList(query entities.MeasurementsQueryContract) ([]entities.Measurement, error)
	GetMeasurementCitiesList(query entities.MeasurementsQueryContract) ([]string, error)
	GetMeasurementLocationsList(query entities.MeasurementsQueryContract) ([]string, error)
	GetMeasurementParametersList(query entities.MeasurementsQueryContract) ([]string, error)
}
