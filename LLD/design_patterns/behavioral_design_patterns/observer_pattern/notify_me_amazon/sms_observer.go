package notifymeamazon

import "fmt"

type SmsObserver struct {
	id int
	o  StockObservable
}

func NewSmsObserver(id int, o StockObservable) *SmsObserver {
	return &SmsObserver{id, o}
}

func (i SmsObserver) GetId() int {
	return i.id
}

func (i SmsObserver) Update() {
	fmt.Printf("Updated sms observer with id :%v with data :%v \n", i.id, i.o.GetData())
}
