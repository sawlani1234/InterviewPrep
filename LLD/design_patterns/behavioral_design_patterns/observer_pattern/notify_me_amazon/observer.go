package notifymeamazon

type NotificationObserver interface {
	GetId() int
	Update()
}
