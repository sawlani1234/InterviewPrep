package user


type User struct {
	Name string 
}

func NewUser(Name string) *User {
	return &User{Name}
}