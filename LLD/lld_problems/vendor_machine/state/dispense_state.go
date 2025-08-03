package state

import (
	"errors"
	"fmt"
	"vendor_machine/enum"
	itemshelf "vendor_machine/item_shelf"
)

type DispenseState struct {
  vendingMachine *VendingMachine 
}

func NewDispenseState(vendingMachine *VendingMachine) *DispenseState {
    return &DispenseState{vendingMachine}
}

func (i *DispenseState) InitState() error {
   return errors.New("incorrect state transition requested")
}

func(i *DispenseState) InsertMoney(money []enum.Money) error {
   return errors.New("incorrect state transition requested")
}

func (i *DispenseState) ChoseProduct(productCode,quantity int) error {
    return errors.New("incorrect state transition requested")
}


func (i *DispenseState) DispenseProduct() error {
    fmt.Println("Dispensing Product ",i.vendingMachine.product.GetItem()," with price ",i.vendingMachine.product.GetPrice())
	
	if i.vendingMachine.GetTotalChange() > 0 {
		fmt.Println("Total Change",i.vendingMachine.GetTotalChange())
	}
	i.vendingMachine.money = i.vendingMachine.money[:0]
	i.vendingMachine.product = itemshelf.ItemShelf{}
	i.vendingMachine.setState(NewInitialState(i.vendingMachine))
	return  nil
}

func (i *DispenseState) Refund() error {
    return errors.New("incorrect state transition requested")
}
