package parkingspacefinderstrategy

import (
	"container/heap"
	"errors"
	"solid_design/parking_lot/entity"
	"solid_design/parking_lot/helper"
)


type NearElevatorStrategy struct {
	nearElevatorHeap *helper.NearElevatorHeap
}

func NewNearElevatorStrategy(parkingSpaces []entity.ParkingSpace) NearElevatorStrategy {

	nearElevatorHeap := &helper.NearElevatorHeap{}
	heap.Init(nearElevatorHeap)

	for _, val := range parkingSpaces {
		heap.Push(nearElevatorHeap, val)
	}

	return NearElevatorStrategy{
		nearElevatorHeap: nearElevatorHeap,
	}
}

func (n NearElevatorStrategy) Find() (entity.ParkingSpace, error) {

	if n.nearElevatorHeap.Len() <= 0 {
		return nil, errors.New("no parking space left")
	}
	return (*n.nearElevatorHeap)[0], nil
}

func (n NearElevatorStrategy) Fill(ps entity.ParkingSpace) error {
	x := heap.Pop(n.nearElevatorHeap).(entity.ParkingSpace)
	if x.GetID() != ps.GetID() {
		return errors.New("parking space cannot be filled")
	}
	return nil
}

func (n NearElevatorStrategy) Empty(ps entity.ParkingSpace) {
	heap.Push(n.nearElevatorHeap, ps)
}
