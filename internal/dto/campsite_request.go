package dto

type CampsiteRequest struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Location      string  `json:"location"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Area          float64 `json:"area"`
	PricePerNight float64 `json:"price_per_night"`
}
