package entity 


type InternalButton struct  {
	liftId int 
	dispatcher InternalButtonDispatcher
}


func NewInternalButton(liftID int,dispatcher InternalButtonDispatcher) InternalButton {
	return InternalButton{liftID,dispatcher}
}



func (i InternalButton) Press(button int) {

	// should have a validation for button is a floor or not 
	i.dispatcher.SubmitRequest(button,i.liftId)
}



