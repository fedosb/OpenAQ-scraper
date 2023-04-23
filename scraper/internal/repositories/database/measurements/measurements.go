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

func (d *database) GetMeasurementsList() ([]entities.Measurement, error) {

	var data []Measurement
	var list []entities.Measurement

	db := d.Model(&Measurement{})

	res := db.Find(&data)
	if res.Error != nil {
		return list, res.Error
	}

	list = buildMeasurementEntities(data)

	return list, nil
}
