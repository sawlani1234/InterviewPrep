package seat

import "bookmyshow/enum"


type Seat struct {
	Id int 
	Number int 
	SeatType enum.SeatType
}

func NewSeat(Id int,Number int,SeatType enum.SeatType) *Seat{
	return &Seat{Id,Number,SeatType}
}