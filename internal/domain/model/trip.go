package model

type Trip struct {
	ID          int32 `json:"id"`
	Origin      *City
	Destination *City
	Dates       string  `json:"dates"`
	Price       float64 `json:"price"`
}
