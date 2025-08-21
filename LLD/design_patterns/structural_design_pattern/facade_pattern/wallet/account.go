package wallet 


type Account struct {
	name string 
}


func NewAccount(name string) *Account {
	return &Account{name}
}


func (a Account) validateAccount(n string) bool  {
	return a.name == n
}

