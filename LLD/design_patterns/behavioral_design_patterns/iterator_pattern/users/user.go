package users


type User struct {
	id int 
	name string 
}

func NewUser(id int,name string) User {
	return User{id,name}
}