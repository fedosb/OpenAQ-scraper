package service

import "TPBDM/scraper/internal/entities"

func (s *service) GetMeasurementsList() ([]entities.Measurement, error) {
	res, err := s.r.DB.Measurements.GetMeasurementsList()
	if err != nil {
		return nil, err
	}

	return res, nil
}
