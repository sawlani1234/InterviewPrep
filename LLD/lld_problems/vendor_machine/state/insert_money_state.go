package state

import (
	"errors"
	"fmt"
	"vendor_machine/enum"
	itemshelf "vendor_machine/item_shelf"
)

type InsertMoneyState struct {
	vendingMachine *VendingMachine
}

func NewInsertMoneyState(vendingMachine *VendingMachine) *InsertMoneyState {
	return &InsertMoneyState{vendingMachine}
}

func (i *InsertMoneyState) InitState() error {
	return errors.New("incorrect state transition requested")
}

func (i *InsertMoneyState) InsertMoney(money []enum.Money) error {

	i.vendingMachine.money = append(i.vendingMachine.money, money...)
	i.vendingMachine.setState(NewChooseProductState(i.vendingMachine))
	return nil

}

func (i *InsertMoneyState) ChoseProduct(productCode, quantity int) error {
	return errors.New("incorrect state transition requested")
}

func (i *InsertMoneyState) DispenseProduct() error {
	return errors.New("incorrect state transition requested")
}

func (i *InsertMoneyState) Refund() error {
	fmt.Println(i.vendingMachine.GetTotalMoney(), " refunded")
	i.vendingMachine.money = i.vendingMachine.money[:0]
	i.vendingMachine.product = itemshelf.ItemShelf{}
	i.vendingMachine.setState(NewInitialState(i.vendingMachine))
	return nil
}
