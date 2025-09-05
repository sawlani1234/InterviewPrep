package manager

import (
	"splitwise_design/splitwise/entities"
	"splitwise_design/splitwise/enums"
)

func NewTransactionFactory(user entities.User, amount int, group *entities.Group, splits []entities.Split, txntype enums.TxnType, splitStrategy entities.SplitStrategy) entities.Transaction {
	switch txntype {
	case enums.EXPENSE:
		{
			return entities.NewExpense(user, group, amount, splits, splitStrategy)
		}
	case enums.SETTLE:
		{
			return entities.NewSettle(user, splits[0], group)
		}
	}

	return nil
}
