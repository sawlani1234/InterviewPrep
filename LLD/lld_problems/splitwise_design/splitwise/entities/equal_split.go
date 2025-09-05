package entities

import "errors"

type EqualSplitStrategy struct {
}

func (e EqualSplitStrategy) Split(amount int, splits []Split) ([]Split, error) {

	firstAmt := splits[0].GetAmount()
	sum := 0
	for i := 1; i < len(splits); i++ {
		sum += splits[i].GetAmount()
		if splits[i].GetAmount() != firstAmt {
			return []Split{}, errors.New("not an equal split")
		}
	}

	if sum != amount {
		return []Split{},errors.New("incorrect split")
	}
	return splits, nil
}