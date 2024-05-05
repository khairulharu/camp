package domain

import "campsite/internal/dto"

type BookingService interface {
	Camp(campID int64) dto.Response
}
