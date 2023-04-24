package service

import (
	"TPBDM/scraper/internal/entities"
	"TPBDM/scraper/internal/repositories"
	"sync"
)

type service struct {
	r             repositories.Repository
	scrapingMutex sync.Mutex
}

type Service interface {
	measurementsMethods
	BeginScraping(contract entities.ScrapingQueryContract) error
}

// New ...
func New(r repositories.Repository) Service {
	return &service{
		r: r,
	}
}
