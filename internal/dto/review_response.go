package dto

import "time"

type ReviewResponse struct {
	ID         int64     `json:"id"`
	CampsiteID int64     `json:"campsite_id"`
	UserID     int64     `json:"user_id"`
	Rating     int8      `json:"rating"`
	Comment    string    `json:"comment"`
	ReviewDate time.Time `json:"riview_date"`
}
