package measurements

import (
	"TPBDM/scraper/internal/entities"
)

func buildMeasurementModel(entity entities.Measurement) MeasurementModel {
	return MeasurementModel{
		ID:        entity.ID,
		DateUTC:   entity.DateUTC,
		Value:     entity.Value,
		Parameter: entity.Parameter,
		Unit:      entity.Unit,
		Country:   entity.Location,
		Location:  entity.Location,
		City:      entity.City,
	}
}

func buildMeasurementEntity(dbModel MeasurementModel) entities.Measurement {
	return entities.Measurement{
		ID:        dbModel.ID,
		DateUTC:   dbModel.DateUTC,
		Value:     dbModel.Value,
		Parameter: dbModel.Parameter,
		Unit:      dbModel.Unit,
		Country:   dbModel.Location,
		Location:  dbModel.Location,
		City:      dbModel.City,
	}
}

func buildMeasurementEntities(dbModels []MeasurementModel) (res []entities.Measurement) {
	for _, dbModel := range dbModels {
		res = append(res, buildMeasurementEntity(dbModel))
	}

	return
}
