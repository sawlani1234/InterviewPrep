package house

import (
	"errors"
	"reflect"
)


type Director struct {
   builder HouseBuilder
}

func NewDirector(builder HouseBuilder) Director {
	return Director{builder}
}


func (d Director) BuildObject() (House,error)  {
	if reflect.TypeOf(d.builder) == reflect.TypeOf(IglooBuilder{}) {
		return d.builder.SetHouseName("igllos-director").SetDoors(3).SetWalls(5).Build(),nil

	} else if reflect.TypeOf(d.builder) == reflect.TypeOf(VillaBuilder{}) {
		return d.builder.SetHouseName("villa-sawlani-director").SetDoors(5).SetWalls(5).Build(),nil
	}
	return House{},errors.New("invalid builder type chosen")
}





