package connector

import "fmt"

type Epson struct {

}

func (e Epson) Print(x string ) {
	fmt.Println("Printing from Epson printer ",x)
}