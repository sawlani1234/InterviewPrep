package users


type UserCollection struct {
	users []User
	iterator Iterator
}

func NewUserCollection(users []User) *UserCollection{
	 return &UserCollection{users,nil}
}


func (u *UserCollection) CreateIterator() Iterator {
	u.iterator = newUserIterator(u.users)
	return u.iterator
}