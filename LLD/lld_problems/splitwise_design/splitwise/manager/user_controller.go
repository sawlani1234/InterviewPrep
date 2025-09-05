package manager

import user "splitwise_design/splitwise/entities"

type UserManager struct {
	users []user.User
}

func NewUserManager() *UserManager{
	return &UserManager{make([]user.User, 0)}
}


func (u *UserManager) Add(user user.User) {
	u.users = append(u.users, user)
}

func (u *UserManager) Get(userId int) user.User {

	for i := 0; i < len(u.users); i++ {
		if u.users[i].GetId() == userId {
			return u.users[i]
		}

	}

	return user.User{}
}
