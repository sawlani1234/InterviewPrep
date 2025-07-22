package helper

import "solid_design/parking_lot/entity"

type NearELevatorAndNearEscalatorHeap []entity.ParkingSpace

func (h *NearELevatorAndNearEscalatorHeap) Len() int {
	return len(*h)
}

func (h *NearELevatorAndNearEscalatorHeap) Less(i, j int) bool {
	return (*h)[i].GetDistancefromElevatorAndEscalator() < (*h)[j].GetDistancefromElevatorAndEscalator()
}

func (h *NearELevatorAndNearEscalatorHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *NearELevatorAndNearEscalatorHeap) Push(p interface{}) {
	*h = append(*h, p.(entity.ParkingSpace))
}

func (h *NearELevatorAndNearEscalatorHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
