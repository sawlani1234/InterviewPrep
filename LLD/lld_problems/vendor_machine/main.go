package main

import (
	// "fmt"
	"vendor_machine/enum"
	itemshelf "vendor_machine/item_shelf"
	"vendor_machine/state"
)

/*
1. Use State Pattern when inner object changes it state 
2. If we do not have states here all logic of managing state would be inside vending machiine struct
which will violate SRP
3. It does not violate ISP because 
	a. there is only one client and it has to call all methods ,if we segregate then it will lead 
	to multiple if else statement by the client which will defat the purpose of state pattern
	b. if clients have different needs then segregate the interface 

*/

func main() {

	itemShelfs := []itemshelf.ItemShelf{
		 *itemshelf.NewItemShelf(enum.CHOCOLATE,3,20,101),
		 *itemshelf.NewItemShelf(enum.COKE,4,30,102),
		 *itemshelf.NewItemShelf(enum.LAYS,3,10,103),
	}

	vendingMachine := state.NewVendingMachine(itemShelfs)
    vendingMachine.GetState().InitState()
    vendingMachine.GetState().InsertMoney([]enum.Money{enum.Twenty,enum.Twenty})
	vendingMachine.GetState().ChoseProduct(101,2)
	vendingMachine.GetState().DispenseProduct()

	
}