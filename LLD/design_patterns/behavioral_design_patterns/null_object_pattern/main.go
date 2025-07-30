package main

import (
	"fmt"
	"vehicle/vehicle"
)


/*
Null object pattern is used in following case

1. When we dont want to have non-null checks because any type of object can be created
based on client demand wherease the case client can send lot of dirty objects 
and its fine for client to have default values rather than error on if cetain objects
are  not present 




*/

func main() {
	
	factory := vehicle.NewVehicleFactory()
	fmt.Println(factory.GetVehicleFactory("CAR").GetTopSpeed())
	fmt.Println(factory.GetVehicleFactory("TRUCK").GetTopSpeed()) // does not panic give default values
}