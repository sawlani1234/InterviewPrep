package paymentstrategy

type CreditCard struct {
}

func (c CreditCard) Pay(amount float64) error {
	return nil
}
