package entity

import (
	"solid_design/elevator_system/enum"
	
)

type ELevatorController struct {
	elevatorMovementStrategy ElevatorMovementStrategy
	elevator Elevator
}

func (e ELevatorController) SubmitNewRequest(floor int, direction enum.Direction) {
     e.elevatorMovementStrategy.AddNewRequest(floor,direction)
}