package entity

import "time"

type Bill struct {
	vehicle Vehicle
	user User 
	fromDate time.Time
	toDate time.Time
	isPaid bool 
	strategy PaymentStrategy
}

func NewBill(vehicle Vehicle,user User,fromDate , toDate time.Time) *Bill {
	return &Bill{vehicle,user,fromDate,toDate,false,nil}
}

func (b *Bill) GenerateBillAmount() int{
	return int(b.toDate.Sub(b.fromDate).Hours()/24)
}

func (b *Bill) SetPaymentStrategy(strategy PaymentStrategy) {
	b.strategy = strategy
}

func (b *Bill) PayBill() error{
	err :=  b.strategy.Pay()
	if err != nil {
		return err
	}
	b.isPaid = true 
	return nil
}

