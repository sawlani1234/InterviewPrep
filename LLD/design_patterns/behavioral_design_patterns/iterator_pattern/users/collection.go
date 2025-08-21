package users


type Collection interface {
	CreateIterator() Iterator
}