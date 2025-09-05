package entities


type User struct {
	id           int
	name         string
	balanceSheet *BalanceSheet
}

func NewUser(id int, name string) *User {
    
	return &User{id, name, NewBalanceSheet()}
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetName() string {
	return  u.name
}

func (u *User) GetBalanceSheet() *BalanceSheet {
	return u.balanceSheet
}
