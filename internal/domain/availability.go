package domain

import "time"

type Avilability struct {
	ID                int64 `gorm:"primaryKey"`
	CampsiteID        int64
	Date              time.Time `gorm:"type:datetime(3)"`
	QuantityAvailable int64
}
