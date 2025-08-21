package users


type Iterator interface {
	HasNext() bool 
	Next() *User
}