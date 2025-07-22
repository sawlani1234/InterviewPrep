package evictionstrategy

import "fmt"

type Lfu struct {

}

func (l Lfu) Evict() {
	 fmt.Println("evicting lfu")
}