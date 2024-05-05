package domain

import (
	"time"

	"gorm.io/gorm"
)

type Campsite struct {
	ID            int64          `gorm:"primaryKey"`
	Name          string         `db:"name"`
	Location      string         `db:"location"`
	Latitude      float64        `db:"latitude"`
	Longitude     float64        `db:"longitude"`
	Area          float64        `db:"area"`
	PricePerNight float64        `db:"price_per_night"`
	CreatedAt     time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt     time.Time      `gorm:"type:datetime(3)"`
	DeletedAt     gorm.DeletedAt `gorm:"type:datetime(3)"`
}

type CampRepository interface {
	FindByID(campID int64) *Campsite
	FindByName(name string) *Campsite
	GetAll() *[]Campsite
}

type CampService interface {
}
