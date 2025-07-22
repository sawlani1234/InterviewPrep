package entity
import "errors"

type PaymentStrategy interface {
   Pay() error

}

type CreditCard struct {

}

func (c CreditCard) Pay() error {
	return nil
}

type UPI struct {
	
}

func (u UPI) Pay() error {
	return nil
}

type DebitCard struct {
	
}

func (d DebitCard) Pay() error {
	return nil
}

func GetPaymentStrategy(id int) (PaymentStrategy,error) {
	switch id {
	case 1:
		return CreditCard{},nil
	case 2:
		return UPI{},nil
	case 3:
		return DebitCard{},nil
	default:
		return nil,errors.New("invalid payment strategy")		

	}
}