package state

import (
	"errors"
	"fmt"
	"vendor_machine/enum"
)

type InitialState struct {
	vendingMachine *VendingMachine
}

func NewInitialState(vendingMachine *VendingMachine) *InitialState {
	return &InitialState{vendingMachine}
}

func (i *InitialState) InitState() error {
	fmt.Println("Initialised vendor machine")

	i.vendingMachine.setState(NewInsertMoneyState(i.vendingMachine))
	return nil
}

func (i *InitialState) InsertMoney(money []enum.Money) error {
	return errors.New("incorrect state transition requested")
}

func (i *InitialState) ChoseProduct(productCode, quantity int) error {
	return errors.New("incorrect state transition requested")
}

func (i *InitialState) DispenseProduct() error {
	return errors.New("incorrect state transition requested")
}

func (i *InitialState) Refund() error {
	return errors.New("incorrect state transition requested")
}
