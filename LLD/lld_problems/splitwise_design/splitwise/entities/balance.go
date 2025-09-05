package entities


type Balance struct {
    owed int 
	given int 
}

func NewBalance() *Balance {
	return &Balance{}
}

func (b *Balance) GetTotalOwed() int {
	return b.owed
}

func (b *Balance) GetTotalGiven() int {
	return  b.given
}