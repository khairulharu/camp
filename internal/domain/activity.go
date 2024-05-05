package domain

import "context"

type Activity struct {
	ID          int64
	Name        string
	Description string `gorm:"type:text"`
}

type ActivityRepository interface {
	FindByID(ctx context.Context, id int64) (Activity, error)
	GetAll(ctx context.Context) ([]Activity, error)
	Insert(ctx context.Context, user *Activity) error
	Update(ctx context.Context, user *Activity) error
}
