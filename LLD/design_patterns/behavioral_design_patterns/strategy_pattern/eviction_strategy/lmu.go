package evictionstrategy

import "fmt"


type Lmu struct {

}

func (l Lmu) Evict() {
	fmt.Println("evictig lmu")
}