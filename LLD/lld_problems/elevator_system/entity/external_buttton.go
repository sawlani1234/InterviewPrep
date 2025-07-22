package entity 

import "solid_design/elevator_system/enum"

type ExternalButton struct {
	floor int 
	dispatcher ExternalButtonDispatcher
}

func NewExternalButton(floor int,dispatcher ExternalButtonDispatcher) ExternalButton {
	return ExternalButton{
		floor: floor,
		dispatcher: dispatcher,
	}
}

func (e ExternalButton) PressButton(direction enum.Direction) {
	e.dispatcher.SubmitRequest(e.floor,direction)
}
	


