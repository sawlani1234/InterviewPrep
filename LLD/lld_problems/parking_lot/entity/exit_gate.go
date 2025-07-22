package entity

type ExitGate struct {
	Vehicle             Vehicle
	Ticket              Ticket
	ParkingSpaceManager ParkingSpaceManager
}

func NewExitGate(ticket Ticket, parkingSpaceManager ParkingSpaceManager) ExitGate {
	return ExitGate{
		Ticket:              ticket,
		ParkingSpaceManager: parkingSpaceManager,
	}
}
