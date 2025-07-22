package house

import "fmt"

type House struct {
	name string 
	doors int 
	walls int 
}

func (h House) Print() {
	fmt.Println ("Hosue baby ",h.name, " with ", h.doors , "  doors and ",h.walls ," walls")
}