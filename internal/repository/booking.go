package repository

import (
	"campsite/internal/domain"
	"context"

	"gorm.io/gorm"
)

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) domain.BookingRepository {
	return &bookingRepository{
		db: db,
	}
}

// FindByID implements domain.ReviewRepository.
func (b *bookingRepository) FindByID(ctx context.Context, id int64) (domain.Booking, error) {
	var booking domain.Booking
	err := b.db.WithContext(ctx).Table("bookings").Where("id = ?", &id).First(&booking).Error
	if err != nil {
		return domain.Booking{}, err
	}
	return booking, nil
}

// GetAll implements domain.bookingRepository.
func (b *bookingRepository) GetAll(ctx context.Context) ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := b.db.WithContext(ctx).Table("bookings").Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

// Insert implements domain.bookingRepository.
func (b *bookingRepository) Insert(ctx context.Context, booking *domain.Booking) error {
	return b.db.WithContext(ctx).Table("bookings").Create(&booking).Error
}

// Update implements domain.bookingRepository.
func (b *bookingRepository) Update(ctx context.Context, booking *domain.Booking) error {
	return b.db.WithContext(ctx).Table("bookings").Model(&domain.Booking{}).Where("id = ?", &booking.ID).Updates(&booking).Error
}

// Delete implements domain.BookingRepository.
func (b *bookingRepository) Delete(ctx context.Context, booking *domain.Booking) error {
	return b.db.WithContext(ctx).Table("bookings").Delete(&booking).Error
}
