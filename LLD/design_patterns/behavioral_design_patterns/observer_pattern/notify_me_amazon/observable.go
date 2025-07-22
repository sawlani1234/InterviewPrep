package notifymeamazon

type StockObservable interface {
	Add(o NotificationObserver)
	Remove(o NotificationObserver)
	Notify()
	SetData(n int)
	GetData() int
}
