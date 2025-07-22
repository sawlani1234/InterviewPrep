package entity

import (
	"solid_design/elevator_system/enum"
	elevatorchoosestrategy "solid_design/elevator_system/strategies/elevator_choose_strategy"
)

type ExternalButtonDispatcher struct {
	elevatorChooseStrategy elevatorchoosestrategy.ElevatorChooseStrategy
	elevatorCoe []ELevatorController
}


func (e ExternalButtonDispatcher) SubmitRequest(floor int,direction enum.Direction) {
	 
	liftID := e.elevatorChooseStrategy.GetElevator(floor,direction)

	for _,val := range e.elevatorCoe {

		if val.elevator.id == liftID {
			val.SubmitNewRequest(floor,val.elevator.display.GetDirection())
			break
		}
	 }
 }

