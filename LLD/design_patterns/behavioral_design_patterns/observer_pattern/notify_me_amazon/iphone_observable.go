package notifymeamazon

import (
	"slices"
)

type IphoneObservable struct {
	data int 
	observers []NotificationObserver
}

func NewIphoneObservable() *IphoneObservable {
	return &IphoneObservable{observers: make([]NotificationObserver, 0)}
}

func (i *IphoneObservable) Add(o NotificationObserver) {
	i.observers = append(i.observers, o)
}

func (i *IphoneObservable) Remove(o NotificationObserver) {

	idx := slices.IndexFunc(i.observers, func(p NotificationObserver) bool {
		return p.GetId() == o.GetId()
	})

	if idx == -1 {
		return
	}

	i.observers = append(i.observers[:idx], i.observers[idx+1:]...)
}

func (i *IphoneObservable) Notify() {

	for j := 0; j < len(i.observers); j++ {
		i.observers[j].Update()
	}
}

func (i *IphoneObservable) GetData() int {
    return i.data
}

func (i *IphoneObservable) SetData(n int) {
    i.data = n
}