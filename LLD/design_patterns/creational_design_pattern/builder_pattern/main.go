package main

import (
	"design_patterns/builder_pattern/house"
	"fmt"
)

/*

1. Builder pattern is creational pattern
2. Used for creating obejcts which have lot of parmatres most of them are operational like configuiratione etc
   this makes constructor very heavy and bad code

3. this problem is sovled by Builder pattern , to make objects step by step on what is absolutely
necessay for the use case

4. Builder pattern also segregates creation of complex of objects from the core object thus adhering to SRP

5. The object created by the pattern is immutable 


Disadvantages 

1. If we add anew field we would have to implement all the implementations of builder 
2. New fields might not be getting used still would be have to be implemented nonethelessa

*/
func main() {
	// without director , directors are not mandatory to create objects they just an abstraction of object
	// creation steps 
    house.NewIglooBuilder().SetHouseName("igloos").SetDoors(5).SetWalls(12).Build().Print()
	house.NewVillaBuilder().SetHouseName("sawlani-villa").SetDoors(3).SetWalls(12).Build().Print()


	// with director
	h1,err := house.NewDirector(house.NewIglooBuilder()).BuildObject()
	if err != nil {
		fmt.Errorf("error creating igloos  house :%v", err.Error())
	}
	fmt.Println("---------------------------------------------")
	h1.Print()
	h2,err := house.NewDirector(house.NewVillaBuilder()).BuildObject()
	if err != nil {
		fmt.Errorf("error creating villa :%v",err.Error())
	}
	h2.Print()
}