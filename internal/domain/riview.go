package domain

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID         int64
	CampsiteID int64
	UserID     int64
	Rating     int8
	Comment    string
	RiviewDate time.Time      `gorm:"type:datetime(3)"`
	CreatedAt  time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt  time.Time      `gorm:"type:datetime(3)"`
	DeletedAt  gorm.DeletedAt `gorm:"type:datetime(3)"`
}
