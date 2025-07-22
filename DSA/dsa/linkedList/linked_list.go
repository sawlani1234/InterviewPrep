package linkedList

type LinkedList struct {
	next *LinkedList
	prev *LinkedList
	val  int
}

func New(val int, prev *LinkedList) *LinkedList {
	return &LinkedList{val: val, prev: prev}
}

func insertInList(l *LinkedList, val int) *LinkedList {
 return &LinkedList{val: val, prev: l, next: l}
}

func traverseList(l *LinkedList) {

}
