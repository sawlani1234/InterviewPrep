package entity

type PaymentStrategy interface {
	Pay(amount float64) error
}