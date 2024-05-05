package dto

import "time"

type ReviewRequest struct {
	ID         int64     `json:"id"`
	CampsiteID int64     `json:"campsite_id"`
	UserID     int64     `json:"user_id"`
	Rating     int8      `json:"rating"`
	Comment    string    `json:"comment"`
	RiviewDate time.Time `json:"riview_date"`
}
