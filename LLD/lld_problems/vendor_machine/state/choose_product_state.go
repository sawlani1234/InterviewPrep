package state

import (
	"errors"
	"fmt"
	"vendor_machine/enum"
	"vendor_machine/item_shelf"
)

type ChooseProductState struct {
  vendingMachine *VendingMachine 
}

func NewChooseProductState(vendingMachine *VendingMachine) *ChooseProductState {
    return &ChooseProductState{vendingMachine}
}

func (i *ChooseProductState) InitState() error {
     return errors.New("incorrect state transition requested")
}

func(i *ChooseProductState) InsertMoney(money []enum.Money) error {
   return errors.New("incorrect state transition requested")
}

func (i *ChooseProductState) ChoseProduct(productCode,quantity int) error {

    for _,itemShelf := range i.vendingMachine.itemshelfs {
		if itemShelf.GetCode() == productCode && itemShelf.GetQuantity() >= quantity {
		    if itemShelf.GetPrice() > i.vendingMachine.GetTotalMoney() {
				return errors.New("less amount added")
			}
			err := itemShelf.SetQuantity(itemShelf.GetQuantity()-quantity)
			i.vendingMachine.product = itemShelf
			if err != nil {
				return err
			}
			break
		}
	}
	i.vendingMachine.setState(NewDispenseState(i.vendingMachine))
	return nil
}


func (i *ChooseProductState) DispenseProduct() error {
    return errors.New("incorrect state transition requested")
}


func (i *ChooseProductState) Refund() error {
	fmt.Println(i.vendingMachine.GetTotalMoney(), " refunded")
	i.vendingMachine.money = i.vendingMachine.money[:0]
    i.vendingMachine.product = itemshelf.ItemShelf{}
    i.vendingMachine.setState(NewInitialState(i.vendingMachine))
	return nil
}
