package entities




type Split struct {
	user User
	amount int 
}

func NewSplit(user User,amount int) Split {
	return  Split{user,amount}
}

func (s Split) GetAmount() int { 
	return s.amount
}

func (s Split) GetUser() User {
	return s.user
}

