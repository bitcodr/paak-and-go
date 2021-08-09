package model

type Trip struct {
	ID            int32   `json:"id"`
	OriginID      int32   `json:"origin_Id"`
	DestinationID int32   `json:"destination-id"`
	Dates         string  `json:"dates"`
	Price         float64 `json:"price"`
}