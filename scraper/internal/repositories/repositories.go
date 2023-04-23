package repositories

import (
	"TPBDM/scraper/config"
	"TPBDM/scraper/internal/repositories/database"
	"TPBDM/scraper/internal/repositories/openaq"
)

type Repository struct {
	DB     *database.Container
	OpenAQ openaq.Repository
}

// New ...
func New(config config.Config) (*Repository, error) {
	db, errDB := database.New(config.Database)
	if errDB != nil {
		return &Repository{}, errDB
	}

	openAQ, errOpenAQ := openaq.New(config.OpenAQ)
	if errOpenAQ != nil {
		return &Repository{}, errOpenAQ
	}

	return &Repository{
		DB:     db,
		OpenAQ: openAQ,
	}, nil
}
