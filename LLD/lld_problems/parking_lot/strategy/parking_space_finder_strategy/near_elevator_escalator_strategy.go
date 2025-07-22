package parkingspacefinderstrategy

import (
	"container/heap"
	"solid_design/parking_lot/helper"
	"solid_design/parking_lot/entity"
)
// TODO: Below implementation is incomplete. Please complete the implementation
type NearElevatorAndNearEscalatorStrategy struct {
	nearELevatorAndNearEscalatorHeap *helper.NearELevatorAndNearEscalatorHeap
}

func NewNearElevatorAndNearEscalatorStrategy(parkingSpaces []entity.ParkingSpace) NearElevatorAndNearEscalatorStrategy {
	nearELevatorAndNearEscalatorHeap := &helper.NearELevatorAndNearEscalatorHeap{}
	heap.Init(nearELevatorAndNearEscalatorHeap)

	for _, val := range parkingSpaces {
		heap.Push(nearELevatorAndNearEscalatorHeap, val)
	}

	return NearElevatorAndNearEscalatorStrategy{
		nearELevatorAndNearEscalatorHeap: nearELevatorAndNearEscalatorHeap,
	}

}
func (n NearElevatorAndNearEscalatorStrategy) Find() (entity.ParkingSpace, error) {

	return heap.Pop(n.nearELevatorAndNearEscalatorHeap).(entity.ParkingSpace), nil
}

func (n NearElevatorAndNearEscalatorStrategy) Fill(ps entity.ParkingSpace) error {
	return nil
}

func (n NearElevatorAndNearEscalatorStrategy) Empty(ps entity.ParkingSpace) {
	
}
