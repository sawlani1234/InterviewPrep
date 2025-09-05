package entities


type Expense struct {
	user  User
	group *Group
	amount int 
	splits []Split
	splitStrategy SplitStrategy
}


func NewExpense(user User,group *Group,amount int,splits []Split,splitStrategy SplitStrategy) Expense {
    s,_ := splitStrategy.Split(amount,splits)

	return Expense{user,group,amount,s,splitStrategy}
}

func (e Expense) GetSplits() []Split {
	return  e.splits
}

func (e Expense) GetAmount() int {
	return e.amount
}

func (e Expense) GetFromUser() User{
	return e.user 
}

func (e Expense) GetToUsers() []User {

	toUsers := []User{}

	for i:=0;i<len(e.splits);i++ {
		toUsers = append(toUsers, e.splits[i].GetUser())
	}
	return toUsers
}