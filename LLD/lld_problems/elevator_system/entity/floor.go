package entity


type Floor struct {
	id int 
	button ExternalButton
}

func NewFloor(id int,dispatcher ExternalButtonDispatcher) Floor{
	return Floor {
		id : id,
		button: NewExternalButton(id,dispatcher),
	}
}