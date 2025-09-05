package splitwise

import (
	"fmt"
	"splitwise_design/splitwise/entities"
	"splitwise_design/splitwise/enums"
	"splitwise_design/splitwise/manager"
)



type Splitwise struct {
	userManager *manager.UserManager
	groupManager *manager.GroupManager
	transactionManager *manager.TransactionManager
}

func NewSplitWise() Splitwise {
	return Splitwise{
		userManager: manager.NewUserManager(),
		groupManager: manager.NewGroupManager(),
		transactionManager: manager.NewTransactionManager(),
	}
}


// Complete this function to simulate splitwise expnese and pay
func (s Splitwise) Demo() {

	user1 := *entities.NewUser(1,"Shubham")
	user2 := *entities.NewUser(2,"Genius Pajji")
	user3 := *entities.NewUser(3,"Mental")
	user4 := *entities.NewUser(4,"Nishi")
	
	s.userManager.Add(user1)
	s.userManager.Add(user2)
	s.userManager.Add(user3)
	s.userManager.Add(user4)


	s.groupManager.Add(*entities.NewGroup(1,[]entities.User{user1,user2,user3,user4}))

	s.transactionManager.ExecuteTransaction(
		user1,[]entities.Split{
			 entities.NewSplit(user2,300),
			 entities.NewSplit(user3,400),
			 entities.NewSplit(user4,200),
		},
		900,
		enums.EXPENSE,
		nil,
		entities.NewUnequalSplitStrategy(),
	)

	// User 1 balance sheet print  
	// User 2 balance sheet print  
	fmt.Println(user1.GetName(), "Total given" ,user1.GetBalanceSheet().GetTotalGiven())
	fmt.Println(user1.GetName(), "Total owed", user1.GetBalanceSheet().GetTotalOwed())

	fmt.Println(user2.GetName(), "Total given" ,user2.GetBalanceSheet().GetTotalGiven())
	fmt.Println(user2.GetName(), "Total owed", user2.GetBalanceSheet().GetTotalOwed())


	fmt.Println(user3.GetName(), "Total given" ,user3.GetBalanceSheet().GetTotalGiven())
	fmt.Println(user3.GetName(), "Total owed", user3.GetBalanceSheet().GetTotalOwed())

	fmt.Println(user4.GetName(), "Total given" ,user4.GetBalanceSheet().GetTotalGiven())
	fmt.Println(user4.GetName(), "Total owed", user4.GetBalanceSheet().GetTotalOwed())


	for user,val := range user1.GetBalanceSheet().GetUserToBalanceMap() {
		fmt.Println(user1.GetName(), " balance with  " , user.GetName(),"  is : ", " total owed is ", val.GetTotalOwed(), " given is", val.GetTotalGiven())
	}



	// USer 2 balance sheet 

}
