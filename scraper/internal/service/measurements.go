package service

import "TPBDM/scraper/internal/entities"

type measurementsMethods interface {
	GetMeasurementsList(query entities.MeasurementsQueryContract) (entities.List[entities.Measurement], error)
	GetMeasurementCitiesList(query entities.MeasurementsQueryContract) (entities.List[string], error)
	GetMeasurementLocationsList(query entities.MeasurementsQueryContract) (entities.List[string], error)
	GetMeasurementParameterList(query entities.MeasurementsQueryContract) (entities.List[string], error)
}

func (s *service) GetMeasurementsList(query entities.MeasurementsQueryContract) (entities.List[entities.Measurement], error) {
	var list entities.List[entities.Measurement]

	res, err := s.r.DB.Measurements.GetMeasurementsList(query)
	if err != nil {
		return list, err
	}

	list.Data = res
	list.Meta.Count = len(res)

	return list, nil
}

func (s *service) GetMeasurementCitiesList(query entities.MeasurementsQueryContract) (entities.List[string], error) {
	var list entities.List[string]

	res, err := s.r.DB.Measurements.GetMeasurementCitiesList(query)
	if err != nil {
		return list, err
	}

	list.Data = res
	list.Meta.Count = len(res)

	return list, nil
}

func (s *service) GetMeasurementLocationsList(query entities.MeasurementsQueryContract) (entities.List[string], error) {
	var list entities.List[string]

	res, err := s.r.DB.Measurements.GetMeasurementLocationsList(query)
	if err != nil {
		return list, err
	}

	list.Data = res
	list.Meta.Count = len(res)

	return list, nil
}

func (s *service) GetMeasurementParameterList(query entities.MeasurementsQueryContract) (entities.List[string], error) {
	var list entities.List[string]

	res, err := s.r.DB.Measurements.GetMeasurementParametersList(query)
	if err != nil {
		return list, err
	}

	list.Data = res
	list.Meta.Count = len(res)

	return list, nil
}
