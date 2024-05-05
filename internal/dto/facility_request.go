package dto

type FacilityRequest struct {
	ID          int64  `json:"id"`
	CampsiteID  int64  `json:"campsite_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
