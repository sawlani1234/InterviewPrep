package entity 


type Building struct {
	floors []Floor
}

func NewBuilding(numberOfFloors int,dispatcher ExternalButtonDispatcher) Building {

    floors := []Floor{}
	
	for i:=0;i<numberOfFloors;i++ {
		floors = append(floors,NewFloor(i,dispatcher))
	}
 
	return Building{floors}
}