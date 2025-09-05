package entities


type Settle struct {
	fromUser User
	split Split
	group *Group
}

func NewSettle(fromUser User,split Split,group *Group) Settle {
	return Settle{fromUser,split,group}
}


func (s Settle) GetSplits() []Split {
	return []Split{s.split}
}

func (s Settle) GetAmount() int {
	return s.split.GetAmount()
}

func (s Settle) GetFromUser() User{
	return s.fromUser
}

func (s Settle) GetToUsers() []User {

	return []User{s.split.GetUser()}
}
