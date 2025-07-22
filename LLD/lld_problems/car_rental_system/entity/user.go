package entity

type User struct {

	id int
	nationalId string
	name string 

}

func NewUser(id int,nationalId ,name string) User{
	return User{id,nationalId,name}
}

func (u User) GetId() int {
	return u.id
}