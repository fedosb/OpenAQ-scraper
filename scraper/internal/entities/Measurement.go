package entities

import (
	"github.com/google/uuid"
	"time"
)

type Measurement struct {
	ID        uuid.UUID `json:"id"`
	DateUTC   time.Time `json:"dateUTC"`
	Value     float32   `json:"value"`
	Parameter string    `json:"parameter"`
	Unit      string    `json:"unit"`
	Country   string    `json:"country"`
	Location  string    `json:"location"`
	City      string    `json:"city"`
}
