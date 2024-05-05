package domain

import (
	"context"
	"time"
)

type Availability struct {
	ID                int64 `gorm:"primaryKey"`
	CampsiteID        int64
	Date              time.Time `gorm:"type:datetime(3)"`
	QuantityAvailable int64
}

type AvailabilityRepository interface {
	FindByID(ctx context.Context, id int64) (Availability, error)
	GetAll(ctx context.Context) ([]Availability, error)
	Insert(ctx context.Context, availability *Availability) error
	Update(ctx context.Context, availability *Availability) error
}
