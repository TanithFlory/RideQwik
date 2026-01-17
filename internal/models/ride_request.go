package models

type RideRequest struct {
	PickupLatitude   float64 `json:"pickup_latitude" binding:"required"`
	PickupLongitude  float64 `json:"pickup_longitude" binding:"required"`
	DropoffLatitude  float64 `json:"dropoff_latitude" binding:"required"`
	DropoffLongitude float64 `json:"dropoff_longitude" binding:"required"`
	PickupAddress    string  `json:"pickup_address"`
	DropoffAddress   string  `json:"dropoff_address"`
	UberToken        string  `json:"uber_token"` // Frontend sends Uber OAuth token
}

type RideOption struct {
	Platform      string  `json:"platform"`
	ProductID     string  `json:"product_id"`
	DisplayName   string  `json:"display_name"`
	EstimatedFare float64 `json:"estimated_fare"`
	ETAMinutes    int     `json:"eta_minutes"`
	Currency      string  `json:"currency"`
}

type RideResponse struct {
	RequestID string       `json:"request_id"`
	Options   []RideOption `json:"options"`
}
