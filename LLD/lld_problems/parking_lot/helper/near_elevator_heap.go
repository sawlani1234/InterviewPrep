package helper

import "solid_design/parking_lot/entity"

type NearElevatorHeap []entity.ParkingSpace

func (h *NearElevatorHeap) Len() int { return len(*h) }

func (h *NearElevatorHeap) Less(i, j int) bool {
	return (*h)[i].GetDistanceFromElevator() < (*h)[j].GetDistanceFromElevator()
}

func (h *NearElevatorHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *NearElevatorHeap) Push(x interface{}) {
	*h = append(*h, x.(entity.ParkingSpace))
}

func (h *NearElevatorHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

