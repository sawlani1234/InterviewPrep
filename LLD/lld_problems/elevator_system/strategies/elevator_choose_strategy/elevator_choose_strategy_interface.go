package elevatorchoosestrategy

import "solid_design/elevator_system/enum"

type ElevatorChooseStrategy interface {
	GetElevator(floor int,direction enum.Direction) int
}
