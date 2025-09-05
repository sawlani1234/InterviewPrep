package entities



type Group struct {
	id int 
	users []User
	userToBalanceSheetMap map[User]*BalanceSheet
}

func NewGroup(id int,user []User) *Group {
	balanceSheetMap  :=  make(map[User]*BalanceSheet,0)

	for i:=0;i<len(user);i++ {
		balanceSheetMap[user[i]] = NewBalanceSheet()
	}
	return &Group{id,user,balanceSheetMap}
}

func (g Group) GetId() int {
	return  g.id
}

func (g Group) UpdateBalanceSheet(user User,splits []Split) {
	
	owed,given := 0,0
	for i:=0;i<len(splits);i++ {
		if splits[i].GetAmount() > 0 {
			given += splits[i].GetAmount()
			g.userToBalanceSheetMap[splits[i].GetUser()].totalOwed += splits[i].GetAmount()
		} else {
			owed += splits[i].GetAmount()
			g.userToBalanceSheetMap[splits[i].GetUser()].totalGiven -= splits[i].GetAmount()
		}
	}
	g.userToBalanceSheetMap[user].totalOwed += owed 
	g.userToBalanceSheetMap[user].totalGiven += given

}