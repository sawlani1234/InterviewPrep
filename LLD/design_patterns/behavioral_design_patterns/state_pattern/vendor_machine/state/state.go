package state

import "vendor_machine/enum"

type State interface {
	InitState() error 
	InsertMoney([]enum.Money) error
	ChoseProduct(productCode ,quantity int) error 
	DispenseProduct() error
	Refund() error 
}