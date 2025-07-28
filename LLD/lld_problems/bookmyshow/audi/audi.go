package audi

import "bookmyshow/seat"

type Audi struct {
	Id int 
	Name string 
	Seats []seat.Seat
}

func NewAudi(Id int,Name string,Seats []seat.Seat) *Audi {
	return &Audi{Id,Name,Seats}
}