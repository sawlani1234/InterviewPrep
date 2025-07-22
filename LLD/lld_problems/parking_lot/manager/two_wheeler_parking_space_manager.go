package manager

import (
	"math/rand"
	"solid_design/parking_lot/entity"
	finderStrategy "solid_design/parking_lot/strategy/parking_space_finder_strategy"
	priceStrategy "solid_design/parking_lot/strategy/price_calculation_strategy"
)

type TwoWheelerParkingSpaceManager struct {
	parkingSpaceManagerBase ParkingSpaceManagerBase
}

func NewTwoWheelerParkingSpaceManager() TwoWheelerParkingSpaceManager {

	ps := make([]*entity.TwoWheelerParkingSpace, 10)
	spaces := []entity.ParkingSpace{}
	for i := 0; i < 10; i++ {

		ps[i] = &entity.TwoWheelerParkingSpace{
			ParkingSpace: entity.ParkingSpaceStruct{
				Id:                               i,
				DistanceFromElevatorAndEscalator: rand.Int() % 97,
				DistanceFromElavator:             rand.Int() % 93,
				PriceCalculationStrategy:         priceStrategy.NewPriceCalculationStrategy("hourly"),
			},
		}
		spaces = append(spaces, ps[i])
	}

	return TwoWheelerParkingSpaceManager{
		parkingSpaceManagerBase: ParkingSpaceManagerBase{
			parkingSpaces:    spaces,
			psFinderStrategy: finderStrategy.NewNearElevatorStrategy(spaces),
		},
	}
}

func (p *TwoWheelerParkingSpaceManager) AddParkingSpace(n int) error {
	return p.parkingSpaceManagerBase.AddParkingSpace(n)
}

func (p *TwoWheelerParkingSpaceManager) RemoveParkingSpace(n int) error {
	return p.parkingSpaceManagerBase.RemoveParkingSpace(n)
}

func (p *TwoWheelerParkingSpaceManager) FindParkingSpace() (entity.ParkingSpace, error) {
	return p.parkingSpaceManagerBase.FindParkingSpace()
}

func (p *TwoWheelerParkingSpaceManager) BookParkingSpace(ps entity.ParkingSpace) error {
	return p.parkingSpaceManagerBase.BookParkingSpace(ps)
}

func (p *TwoWheelerParkingSpaceManager) GetParkingSpaces() []entity.ParkingSpace {
	return p.parkingSpaceManagerBase.GetParkingSpaces()
}

func (p *TwoWheelerParkingSpaceManager) EmptyParkingSpace(ps entity.ParkingSpace) {
	p.parkingSpaceManagerBase.EmptyParkingSpace(ps)
}
