package measurements

import (
	"TPBDM/scraper/internal/entities"
	"errors"
	"gorm.io/gorm"
)

// database ...
type database struct {
	*gorm.DB
}

func NewMeasurementsRepository(db *gorm.DB) Repository {
	return &database{DB: db}
}

func (d *database) CreateMeasurement(measurement entities.Measurement) error {
	model := buildMeasurementModel(measurement)

	res := d.Create(&model)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("MEASUREMENT WASN'T CREATED")
	}

	return nil
}

func (d *database) GetMeasurementsList(query entities.MeasurementsQueryContract) ([]entities.Measurement, error) {

	var data []Measurement
	var list []entities.Measurement

	db := d.Model(&Measurement{})
	db = d.applyQueryToMeasurements(db, query)

	res := db.Find(&data)
	if res.Error != nil {
		return list, res.Error
	}

	list = buildMeasurementEntities(data)

	return list, nil
}

func (d *database) GetMeasurementCitiesList(query entities.MeasurementsQueryContract) ([]string, error) {

	var data []string

	db := d.Model(&Measurement{})
	db = d.applyQueryToMeasurements(db, query)

	res := db.Distinct().Pluck("city", &data)
	if res.Error != nil {
		return data, res.Error
	}

	return data, nil
}

func (d *database) GetMeasurementLocationsList(query entities.MeasurementsQueryContract) ([]string, error) {

	var data []string

	db := d.Model(&Measurement{})
	db = d.applyQueryToMeasurements(db, query)

	res := db.Distinct().Pluck("location", &data)
	if res.Error != nil {
		return data, res.Error
	}

	return data, nil
}

func (d *database) GetMeasurementParametersList(query entities.MeasurementsQueryContract) ([]string, error) {

	var data []string

	db := d.Model(&Measurement{})
	db = d.applyQueryToMeasurements(db, query)

	res := db.Distinct().Pluck("parameter", &data)
	if res.Error != nil {
		return data, res.Error
	}

	return data, nil
}

func (d *database) applyQueryToMeasurements(db *gorm.DB, query entities.MeasurementsQueryContract) *gorm.DB {
	if query.Parameter != "" {
		db = db.Where("parameter = ?", query.Parameter)
	}

	if query.City != "" {
		db = db.Where("city = ?", query.City)
	}

	if query.Country != "" {
		db = db.Where("country = ?", query.Country)
	}

	if query.Location != "" {
		db = db.Where("location = ?", query.Location)
	}

	return db
}
