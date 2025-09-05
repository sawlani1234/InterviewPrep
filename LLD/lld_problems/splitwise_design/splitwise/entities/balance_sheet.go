package entities

import "splitwise_design/splitwise/enums"



	type BalanceSheet struct {

	user User
	totalOwed int 
	totalGiven int 
	userToBalanceMap map[User]*Balance 
}


func NewBalanceSheet() *BalanceSheet{
	return &BalanceSheet{userToBalanceMap: make(map[User]*Balance)}
}

// add update logic here 

func (b *BalanceSheet) UpdateBalanceSheet(splits []Split,txnType enums.TxnType) {
	// update all balances of current user
	owed,given := 0,0
	for i:=0;i<len(splits);i++ {

	   if txnType == enums.EXPENSE {
		  given+=splits[i].GetAmount()

		  if _, ok := b.userToBalanceMap[splits[i].GetUser()] ; !ok {
			b.userToBalanceMap[splits[i].user] = NewBalance()
		  }
		  b.userToBalanceMap[splits[i].GetUser()].given += splits[i].GetAmount()
	   } else {
		  // implement for settle 
	   }
	   
	}

	b.totalGiven += given
	b.totalOwed += owed

	// update balances of the all users in the split 

	for i:=0;i<len(splits);i++ {
		 balanceSheet := splits[i].user.balanceSheet
		

		 if txnType == enums.EXPENSE {
			if _, ok := b.userToBalanceMap[b.user] ; !ok {
			balanceSheet.userToBalanceMap[b.user] = NewBalance()
		  }
			balanceSheet.userToBalanceMap[b.user].owed += splits[i].amount
			balanceSheet.totalOwed += splits[i].GetAmount()
		 } else {
			// implement for settle 
		 }
	}


}


func (b *BalanceSheet) GetTotalOwed() int {
	return b.totalOwed
}

func (b *BalanceSheet) GetTotalGiven() int {
	return b.totalGiven
}


func (b *BalanceSheet) GetUserToBalanceMap() map[User]*Balance{
	return b.userToBalanceMap
}



