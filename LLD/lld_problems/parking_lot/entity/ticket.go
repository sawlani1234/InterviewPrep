package entity

import "time"

type Ticket struct {
	Id           int
	entryTime    time.Time
	Vehicle      Vehicle
	ParkingSpace ParkingSpace
}

func NewTicket(id int,vehicle Vehicle, parkingSpace ParkingSpace) Ticket {
	return Ticket{
		Id:           id,
		entryTime:    time.Now(),
		Vehicle:      vehicle,
		ParkingSpace: parkingSpace,
	}
}

func (t Ticket) GetEntryTime() time.Time {
	return t.entryTime
}
