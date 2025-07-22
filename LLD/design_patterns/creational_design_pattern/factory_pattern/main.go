package main

import (
	"factory_pattern/shape"
	"fmt"
)

/**

1. Factor pattern is creational pattern used to create objects when there are multiple of objects of a same type
and if client using those objects do not have factory pattern they would have to add nasty if else logic
to get the object , thus in factory pattern it provides a centralise location to create the objects 
by clients , thus clients request for the object and do there operation 
prevents duplicate and messy logic 

**/

func main() {
	shapeFactory := shape.NewShapeFactory()
	sq,_ := shapeFactory.GetShapeFactory("Square", 2, 2)
	fmt.Println(sq.ComputeArea())
	
	rec,_ := shapeFactory.GetShapeFactory("Rectangle",3,4)
	fmt.Println(rec.ComputeArea())

	_,err := shapeFactory.GetShapeFactory("asa",1,2)
	fmt.Println(err)

}
