package domain

type Activity struct {
	ID          int64
	Name        string
	Description string `gorm:"type:text"`
}
