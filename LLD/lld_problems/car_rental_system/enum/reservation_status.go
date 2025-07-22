package enum


type ReservationStatus int 


const (
	NO_RESERVATION ReservationStatus = iota
	SCHEDULED
	SUBMITTED 
	COMPLETED
)