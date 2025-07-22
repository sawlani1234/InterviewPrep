package evictionstrategy

import "fmt"


type Lru struct {

}

func (l Lru) Evict() {
	fmt.Println ("evicting lru")
}