// TODO : complete implementation
package manager
import "solid_design/parking_lot/entity"

type FourWheelerParkingSpaceManager struct {
	parkingSpaceManagerBase ParkingSpaceManagerBase
}

func NewFourWheelerParkingSpaceManager() FourWheelerParkingSpaceManager {
	return FourWheelerParkingSpaceManager{}
}

func (p *FourWheelerParkingSpaceManager) AddParkingSpace(n int) error {
     return p.parkingSpaceManagerBase.AddParkingSpace(n)
}

func (p *FourWheelerParkingSpaceManager) RemoveParkingSpace(n int) error {
	return p.parkingSpaceManagerBase.RemoveParkingSpace(n)
}

func (p *FourWheelerParkingSpaceManager) FindParkingSpace() (entity.ParkingSpace, error) {
	return p.parkingSpaceManagerBase.FindParkingSpace()
}

func (p *FourWheelerParkingSpaceManager) BookParkingSpace(ps entity.ParkingSpace) error {
	return p.parkingSpaceManagerBase.BookParkingSpace(ps)
}

func (p *FourWheelerParkingSpaceManager) GetParkingSpaces() []entity.ParkingSpace {
	return p.parkingSpaceManagerBase.GetParkingSpaces()
}

func (p *FourWheelerParkingSpaceManager) EmptyParkingSpace(ps entity.ParkingSpace) {
	p.parkingSpaceManagerBase.EmptyParkingSpace(ps)
}