package paymentstrategy

type Cash struct {
}

func (c Cash) Pay(amount float64) error {
	return nil
}