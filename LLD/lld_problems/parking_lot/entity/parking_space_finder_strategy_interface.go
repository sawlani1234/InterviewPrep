package entity

type ParkingSpaceFinderStrategy interface {
	Find() (ParkingSpace, error)
	Fill(ps ParkingSpace) error
	Empty(ps ParkingSpace)
}