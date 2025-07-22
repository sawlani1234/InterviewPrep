package entity

import (
	"solid_design/parking_lot/enums"
)


type ParkingSpace interface {
    GetID() int 
	IsOccupied() bool 
	GetVehicle() Vehicle
	SetOccupied()
	Empty()
	GetDistanceFromElevator() int
	GetDistancefromElevatorAndEscalator() int
	GetPriceCalculationStrategy() PriceCalculationStrategy
}


type ParkingSpaceStruct struct {
	Id                               int
	isOccupied                       bool
	vehicle                          Vehicle
	DistanceFromElavator             int
	DistanceFromElevatorAndEscalator int
	vehicleType                      enums.VehicleType
	PriceCalculationStrategy         PriceCalculationStrategy
}

func (p *ParkingSpaceStruct) GetID() int {
	return p.Id
}

func (p *ParkingSpaceStruct) IsOccupied() bool {
	return p.isOccupied
}

func (p *ParkingSpaceStruct) GetVehicle() Vehicle {
	return p.vehicle
}

func (p *ParkingSpaceStruct) SetOccupied() {
	p.isOccupied = true 
}

func (p *ParkingSpaceStruct) Empty() {
	p.isOccupied = false 
}

func (p *ParkingSpaceStruct) GetDistanceFromElevator() int {
	return p.DistanceFromElavator
}

func (p *ParkingSpaceStruct) GetDistancefromElevatorAndEscalator() int {
	return p.DistanceFromElevatorAndEscalator
}

func (p *ParkingSpaceStruct) GetPriceCalculationStrategy() PriceCalculationStrategy {
	return p.PriceCalculationStrategy
}




type TwoWheelerParkingSpace struct {
	ParkingSpace ParkingSpaceStruct
	// some extra variables can be added  for two wheeler
}

func NewTwoWheelerParkingSpace(priceStrategy PriceCalculationStrategy) TwoWheelerParkingSpace {
	return TwoWheelerParkingSpace{
		ParkingSpace: ParkingSpaceStruct{
			isOccupied:               false,
			vehicleType:              enums.TwoWheeler,
			PriceCalculationStrategy: priceStrategy,
		},
	}
}

func (p *TwoWheelerParkingSpace) GetID() int {
	return p.ParkingSpace.Id
}

func (p *TwoWheelerParkingSpace) IsOccupied() bool {
	return p.ParkingSpace.IsOccupied()
}

func (p *TwoWheelerParkingSpace) GetVehicle() Vehicle {
	return p.ParkingSpace.vehicle
}

func (p *TwoWheelerParkingSpace) SetOccupied() {
	 p.ParkingSpace.SetOccupied()
}

func (p *TwoWheelerParkingSpace) Empty() {
	 p.ParkingSpace.Empty()
}

func (p *TwoWheelerParkingSpace) GetDistanceFromElevator() int {
	return p.ParkingSpace.GetDistanceFromElevator()
}

func (p *TwoWheelerParkingSpace) GetDistancefromElevatorAndEscalator() int {
	return p.ParkingSpace.GetDistancefromElevatorAndEscalator()
}

func (p *TwoWheelerParkingSpace) GetPriceCalculationStrategy() PriceCalculationStrategy {
	return p.ParkingSpace.GetPriceCalculationStrategy()
}




type FourWheelerParkingSpace struct {
	ParkingSpace ParkingSpaceStruct
	// some extra variables can be added for four wheeler
}

func NewFourWheelerParkingSpace(priceStrategy PriceCalculationStrategy) FourWheelerParkingSpace {
	return FourWheelerParkingSpace{
		ParkingSpace: ParkingSpaceStruct{
			isOccupied: false,
			vehicleType: enums.FourWheeler, 
			PriceCalculationStrategy: priceStrategy,
		},
	}
}
