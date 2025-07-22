package dequeue

type Dequeue struct {
	head *LinkedList
	tail *LinkedList
}

type LinkedList struct {
	val  int
	next *LinkedList
	prev *LinkedList
}

func NewLinkedList(p int) *LinkedList {
	return &LinkedList{
		val: p,
	}
}

func (q *Dequeue) PushBack(p int) {
	if q.head == nil {
		q.head = NewLinkedList(p)
		q.tail = q.head
	} else {
		q.tail.next = NewLinkedList(p)
		q.tail.next.prev = q.tail
		q.tail = q.tail.next
	}
}
func (q *Dequeue) PopFront() int {
	if q.IsEmpty() {
		return -1
	}
	t := q.head.val
	q.head = q.head.next
	if q.head != nil {
		q.head.prev = nil
	} else {
		q.tail = nil
	}
	return t
}

func (q *Dequeue) PushFront(p int) {
	if q.head == nil {
		q.head = NewLinkedList(p)
		q.tail = q.head
	} else {
		temp := q.head
		q.head = NewLinkedList(p)
		q.head.next = temp
		q.head.next.prev = q.head
	}
}

func (q *Dequeue) PopBack() int {
	if q.IsEmpty() {
		return -1
	}
	val := q.tail.val
	if q.tail == q.head {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.prev
		q.tail.next = nil
	}
	return val
}

func (q *Dequeue) IsEmpty() bool {
	return q.head == nil
}

func (q *Dequeue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.head.val
}

func (q *Dequeue) Back() int {
	if q.IsEmpty() {
		return -1
	}
	return q.tail.val
}

func NewDequeue() *Dequeue {
	return &Dequeue{}
}
