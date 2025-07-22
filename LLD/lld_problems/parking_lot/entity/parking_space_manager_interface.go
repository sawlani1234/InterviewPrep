package entity

type ParkingSpaceManager interface {
	AddParkingSpace(n int) error
	RemoveParkingSpace(n int) error
	BookParkingSpace(ps ParkingSpace) error
	FindParkingSpace() (ParkingSpace, error)
	EmptyParkingSpace(ps ParkingSpace)
}