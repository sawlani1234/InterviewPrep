package itemshelf

import (
	"errors"
	"vendor_machine/enum"
)

type ItemShelf struct {
	item	enum.Item
	quantity int 
	price int 
	code int 
}


func NewItemShelf(item enum.Item,quantity int,price int,code int) *ItemShelf {
	return &ItemShelf{item,quantity,price,code}
}

func (i *ItemShelf) GetItem() enum.Item {
	return i.item
}

func (i *ItemShelf) GetQuantity() int {
	return i.quantity
}


func (i *ItemShelf) GetPrice() int {
	return  i.price
}

func (i *ItemShelf) SetQuantity(q int) error {

	if q < 0 {
		return errors.New("invalid quanitiy selected")
	}
	i.quantity = q
	return nil
}

func (i *ItemShelf) GetCode() int {
	return i.code
}