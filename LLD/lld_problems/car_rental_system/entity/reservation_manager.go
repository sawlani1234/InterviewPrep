package entity

import (
	"fmt"
)

type ReservationManager struct {
    reservations []Reservation
}

func NewReservationManager() *ReservationManager{
	return &ReservationManager{
		reservations: make([]Reservation, 0),
	}
}

func (r *ReservationManager) Add(reservation Reservation) {
	r.reservations = append(r.reservations, reservation)
}

func (r *ReservationManager) Remove(reservation Reservation) error {
	idx := -1
	
	for i:=0;i<len(r.reservations);i++ {
		if r.reservations[i].GetId() == reservation.GetId() {
			idx = i 
			break
		}
	}

	if idx == -1 {
		return fmt.Errorf("no such vehicle present with id:%v",reservation.GetId())

	}
	return nil 
}

func (r *ReservationManager) Update(reservation Reservation) error {
	for i:=0;i<len(r.reservations);i++ {
		if r.reservations[i].GetId() == reservation.GetId() {
			r.reservations[i] = reservation
			return nil 
		}
	}
	return fmt.Errorf("no such vehicle present with id :%v",reservation.GetId())
}