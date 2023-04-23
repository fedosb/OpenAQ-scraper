package service

import (
	"TPBDM/scraper/internal/entities"
	"TPBDM/scraper/internal/repositories"
)

type service struct {
	r repositories.Repository
}

type Service interface {
	GetMeasurementsList() ([]entities.Measurement, error)
}

// New ...
func New(r repositories.Repository) Service {
	return &service{
		r,
	}
}
