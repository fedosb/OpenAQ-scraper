package openaq

import "TPBDM/scraper/internal/entities"

func buildMeasurementEntity(contract MeasurementContract) entities.Measurement {
	return entities.Measurement{
		DateUTC:   contract.Date.Utc,
		Value:     contract.Value,
		Parameter: contract.Parameter,
		Unit:      contract.Unit,
		Country:   contract.Location,
		Location:  contract.Location,
		City:      contract.City,
	}
}

func buildMeasurementEntities(contracts []MeasurementContract) (res []entities.Measurement) {
	for _, dbModel := range contracts {
		res = append(res, buildMeasurementEntity(dbModel))
	}

	return
}
