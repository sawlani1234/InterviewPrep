package queue

/*
Queue implementation using linked list
*/
type Pair struct {
	x int
	y int
}
type Queue struct {
	head *LinkedList
	tail *LinkedList
}

type LinkedList struct {
	coord Pair
	next  *LinkedList
}

func NewLinkedList(p Pair) *LinkedList {
	return &LinkedList{
		coord: p,
	}
}

func (q *Queue) Push(p Pair) {
	if q.head == nil {
		q.head = NewLinkedList(p)
		q.tail = q.head
	} else {
		q.tail.next = NewLinkedList(p)
		q.tail = q.tail.next
	}
}
func (q *Queue) Pop() Pair {
	if q.IsEmpty() {
		return Pair{}
	}
	t := q.head.coord
	q.head = q.head.next
	return t
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Top() Pair {
	if q.IsEmpty() {
		return Pair{}
	}
	return q.head.coord
}

func NewQueue() *Queue {
	return &Queue{}
}

