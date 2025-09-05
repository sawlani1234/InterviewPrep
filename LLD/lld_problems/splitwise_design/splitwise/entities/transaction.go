package entities



type Transaction interface{
	GetSplits() []Split
	GetAmount() int 
	GetFromUser() User
	GetToUsers() []User
}
