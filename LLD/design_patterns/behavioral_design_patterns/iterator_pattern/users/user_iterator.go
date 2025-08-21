package users


type UserIterator struct {
	curIdx int 
	users []User
}

func newUserIterator(users []User) *UserIterator {
	return &UserIterator{0,users}
}

func (u *UserIterator) HasNext() bool {
	return u.curIdx < (len(u.users))
}


func (u *UserIterator) Next() *User {
	x := u.users[u.curIdx]
	u.curIdx++
	return &x
}