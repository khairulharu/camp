package dto

import "time"

type BookingResponse struct {
	ID             int64     `json:"id"`
	CampsiteID     int64     `json:"campsite_id"`
	UserID         int64     `json:"user_id"`
	CheckInDate    time.Time `json:"checkin_date"`
	CheckOutDate   time.Time `json:"chechout_date"`
	NumberOfPeople int64     `json:"number_of_people"`
	TotalCost      float64   `json:"total_cost"`
	Status         int8      `json:"status"`
}
