package repositories

import (
	"TPBDM/scraper/config"
	"TPBDM/scraper/internal/repositories/database"
)

type Repository struct {
	DB *database.Container
}

// New ...
func New(config config.Config) (*Repository, error) {
	db, errDB := database.New(config.Database)
	if errDB != nil {
		return &Repository{}, errDB
	}

	return &Repository{
		DB: db,
	}, nil
}
