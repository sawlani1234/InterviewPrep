package entity

import (
	"solid_design/car_rental/enum"
	"time"
)

type Reservation struct {
	id int
	user User
	vechile Vehicle
	status enum.ReservationStatus
	fromDate time.Time
	toDate time.Time
	Bill *Bill
}

func NewReservation(user User,vehicle Vehicle,status enum.ReservationStatus,fromDate time.Time,toDate time.Time) *Reservation {
	return &Reservation{
		user: user,
		vechile: vehicle,
		status: status,
		fromDate: fromDate,
		toDate: toDate,
		Bill: NewBill(vehicle,user,fromDate,toDate),
	}
}

func (r *Reservation) SetStatus(status enum.ReservationStatus) {
	r.status = status
}

func (r *Reservation) GetStatus() enum.ReservationStatus {
	return r.status
}

func (r *Reservation) GetId() int {
	return r.id
}