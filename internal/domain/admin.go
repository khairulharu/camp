package domain

type Admin struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}
