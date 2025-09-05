package manager

import (
	"splitwise_design/splitwise/entities"
	"splitwise_design/splitwise/enums"
)

type TransactionManager struct {
}

func NewTransactionManager() *TransactionManager {
	return &TransactionManager{}
}

func (t *TransactionManager) ExecuteTransaction(user entities.User, splits []entities.Split, amount int, txnType enums.TxnType,group *entities.Group,splitStrategy entities.SplitStrategy) {

	txn := NewTransactionFactory(user,amount,group,splits,txnType,splitStrategy)

	if group == nil {
		user.GetBalanceSheet().UpdateBalanceSheet(txn.GetSplits(),txnType)
	} else {
		group.UpdateBalanceSheet(user,txn.GetSplits())
	}
}
