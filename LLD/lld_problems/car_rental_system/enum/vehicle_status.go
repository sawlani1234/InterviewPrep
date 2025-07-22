package enum 

type VehicleStatus int 

const (
	INVALID VehicleStatus = iota
	ACTIVE
	INACTIVE 
	BOOKED
)