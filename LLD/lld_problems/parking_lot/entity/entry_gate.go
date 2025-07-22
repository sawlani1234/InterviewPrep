package entity

type EntryGate struct {
	Ticket              Ticket
	Vehicle             Vehicle
	ParkingSpaceManager ParkingSpaceManager
}

func NewEntryGate(vehicle Vehicle, parkingSpaceManager ParkingSpaceManager) EntryGate {
	return EntryGate{
		Vehicle:             vehicle,
		ParkingSpaceManager: parkingSpaceManager,
	}
}
