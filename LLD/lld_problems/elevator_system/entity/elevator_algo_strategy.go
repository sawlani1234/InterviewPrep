package entity

import "solid_design/elevator_system/enum"

type ElevatorMovementStrategy interface {
	AddNewRequest(floor int, direction enum.Direction)
}

