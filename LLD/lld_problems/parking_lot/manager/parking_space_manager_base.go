package manager

import (
	"errors"
	"solid_design/parking_lot/entity"
	
)

type ParkingSpaceManagerBase struct {
    parkingSpaces  []entity.ParkingSpace
	psFinderStrategy entity.ParkingSpaceFinderStrategy
}

func (p *ParkingSpaceManagerBase) AddParkingSpace(n int) error {
     if n <= 0 {
		return errors.New("number of parking spaces should be greater than 0")
	}

	p.parkingSpaces = append(p.parkingSpaces, make([]entity.ParkingSpace, n)...)
	return nil
}

func (p *ParkingSpaceManagerBase) RemoveParkingSpace(n int) error {
	if n <= 0 {
		return errors.New("number of parking spaces should be greater than 0")
	}

	if n > len(p.parkingSpaces) {
		return errors.New("number of parking spaces to remove is greater than available parking spaces")
	}

	p.parkingSpaces = p.parkingSpaces[:len(p.parkingSpaces)-n]
	return nil
}

func (p *ParkingSpaceManagerBase) FindParkingSpace() (entity.ParkingSpace, error) {
	return p.psFinderStrategy.Find()
}

func (p *ParkingSpaceManagerBase) BookParkingSpace(ps entity.ParkingSpace) error {

	if ps.IsOccupied() {
		return errors.New("parking space is already occupied")
	}

	ps.SetOccupied()

	for i, val := range p.parkingSpaces {
		if val.GetID() == ps.GetID() {
			p.parkingSpaces[i] = ps
			break
		}
	}
	return p.psFinderStrategy.Fill(ps)

}

func (p *ParkingSpaceManagerBase) GetParkingSpaces() []entity.ParkingSpace {
	return p.parkingSpaces
}

func (p *ParkingSpaceManagerBase) EmptyParkingSpace(ps entity.ParkingSpace) {
	ps.Empty()
	p.parkingSpaces[ps.GetID()] = ps
	p.psFinderStrategy.Empty(ps)
}
