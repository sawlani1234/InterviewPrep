package state

import (
	"vendor_machine/enum"
	itemshelf "vendor_machine/item_shelf"
)

type VendingMachine struct {
	state      State
	money      []enum.Money
	itemshelfs []itemshelf.ItemShelf
	product    itemshelf.ItemShelf
}

func NewVendingMachine(itemshelfs []itemshelf.ItemShelf) *VendingMachine {

	v := &VendingMachine{itemshelfs: itemshelfs}
	v.state = NewInitialState(v)
	return v
}

func (v *VendingMachine) GetState() State {
	return v.state
}


func (v *VendingMachine) setState(s State) {
	v.state = s
}

func (v *VendingMachine) GetTotalMoney() int {

	sum := 0
	for _, money := range v.money {
		sum += money.Int()
	}

	return sum
}

func (v *VendingMachine) GetTotalChange() int {
	return v.GetTotalMoney() - v.product.GetPrice()
}
