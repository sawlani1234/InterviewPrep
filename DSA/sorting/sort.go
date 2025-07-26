package sorting

type SortingStrategy interface {
	Sort(a []int) []int
}

type Sorter struct {
	Strategy SortingStrategy
}

func (s *Sorter) SetSortingStrategy(strategy SortingStrategy) {
	s.Strategy = strategy
}
