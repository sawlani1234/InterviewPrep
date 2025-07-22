package notifymeamazon

import "fmt"

type InboxObserver struct {
	id int
	o  StockObservable
}

func NewInboxObserver(id int, o StockObservable) *InboxObserver {
	return &InboxObserver{id, o}
}

func (i InboxObserver) GetId() int {
	return i.id
}

func (i InboxObserver) Update() {
	fmt.Printf("Updated inbox observer with id :%v with data :%v \n", i.id, i.o.GetData())
}
