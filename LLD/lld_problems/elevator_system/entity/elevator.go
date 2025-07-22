package entity

import "fmt"


type Elevator struct {
	id int 
	display Display
	door Door
	button InternalButton
	
}

func NewElevator(id int,dispatcher InternalButtonDispatcher) Elevator{
	return Elevator{
		id,
		*NewDisplay(),
		NewDoor(),
		NewInternalButton(id,dispatcher),
	}
}

func (e Elevator) PrintElevator() {
	fmt.Println("Elavator :",e.id, "currently at floor :",e.display.GetFloor()," with status :",e.display.GetStatus()," with direction :",e.display.GetDirection())
}



