package stack

type Stack struct {
	val []rune
}

func (s *Stack) Push(input rune) {
	s.val = append(s.val, input)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return -1
	}
	ans := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return ans
}

func (s *Stack) Top() rune {
	if s.IsEmpty() {
		return -1
	}
	return s.val[len(s.val)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.val) == 0
}

func NewStack() *Stack {
	return &Stack{val: make([]rune, 0)}
}
