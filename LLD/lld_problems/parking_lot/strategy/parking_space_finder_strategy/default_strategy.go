package parkingspacefinderstrategy

import "solid_design/parking_lot/entity"

type Default struct {
}

func NewDefault() Default {
	return Default{}
}

func (d Default) Find() (entity.ParkingSpace, error) {
	return nil, nil
}

func (d Default) Fill(ps entity.ParkingSpace) error {
	return nil

}

func (d Default) Empty(ps entity.ParkingSpace) {
	
}
