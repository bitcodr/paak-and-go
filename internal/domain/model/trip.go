package model

type Trip struct {
	ID            int32
	OriginId      int32
	DestinationId int32
	Dates         string
	Price         float64
}


//
//var trips = []trip{
//	{id: 1, originId: 1, destinationId: 2, dates: "Mon Tue Wed Fri", price: 40.55},
//	{id: 2, originId: 2, destinationId: 1, dates: "Sat Sun", price: 40.55},
//	{id: 3, originId: 3, destinationId: 6, dates: "Mon Tue Wed Thu Fri", price: 32.10},
//}
