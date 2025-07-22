package main

import (
	"fmt"
	notifymeamazon "observer_pattern/notify_me_amazon"
)

/*

1. Design Notify Me
2. Observer pattern has observer who observe observable
3. Any state change in Observable notifies all observer
4.
*/

func main() {

	iphoneObservable := notifymeamazon.NewIphoneObservable()

	shubahmObserver := notifymeamazon.NewInboxObserver(1, iphoneObservable)
	nishiObserver := notifymeamazon.NewInboxObserver(2, iphoneObservable)
	deepPawanObserver := notifymeamazon.NewInboxObserver(3, iphoneObservable)
	shwetaObserver := notifymeamazon.NewInboxObserver(4, iphoneObservable)

	shubhamSmsObserver := notifymeamazon.NewSmsObserver(3,iphoneObservable)
	nishiSmsObserver := notifymeamazon.NewSmsObserver(4,iphoneObservable)

	iphoneObservable.Add(*shubahmObserver)
	iphoneObservable.Add(*nishiObserver)
	iphoneObservable.Add(*deepPawanObserver)
	iphoneObservable.Add(*shwetaObserver)
	iphoneObservable.Add(*shubhamSmsObserver)
	iphoneObservable.Add(*nishiSmsObserver)
   

	fmt.Println("Starting to notify")
	iphoneObservable.SetData(3)
	iphoneObservable.Notify()

	iphoneObservable.Remove(*nishiObserver)

	fmt.Println("\n Starting to Notify")

	iphoneObservable.Notify()

}
