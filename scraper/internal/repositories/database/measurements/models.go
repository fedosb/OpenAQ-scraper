package measurements

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MeasurementModel struct {
	*gorm.Model

	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	DateUTC   time.Time `gorm:"type:timestamptz;not null"`
	Value     float32   `gorm:"type:real;not null"`
	Parameter string    `gorm:"type:varchar;not null"`
	Unit      string    `gorm:"type:varchar;not null"`
	Country   string    `gorm:"type:varchar;not null;index"`
	Location  string    `gorm:"type:varchar;not null;index"`
	City      string    `gorm:"type:varchar;not null;index"`
}

func (m *MeasurementModel) BeforeCreate(_ *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}

	return nil
}
