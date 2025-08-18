package connector

import "fmt"


type Hp struct {

}

func (h Hp) Print(x string) {
	fmt.Println("Printing from HP printer", x)
}