package entity

type InternalButtonDispatcher struct {
	elevatorCOE []ELevatorController
}


func NewInternalButtonDispatcher(elevatorCOE []ELevatorController) InternalButtonDispatcher{
	return InternalButtonDispatcher{elevatorCOE}

}

 
 func (i InternalButtonDispatcher) SubmitRequest(floor int,liftID int) {
	 
	for _,val := range i.elevatorCOE {

		if val.elevator.id == liftID {
			val.SubmitNewRequest(floor,val.elevator.display.GetDirection())
			break
		}
	 }
 }