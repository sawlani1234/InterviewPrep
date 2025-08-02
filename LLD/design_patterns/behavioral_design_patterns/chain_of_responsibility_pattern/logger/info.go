package logger

import (
	"fmt"
	"solid_design/chain_of_responsibility/enum"
)

type Info struct {
	next Logger
}

func (i *Info) Log(logLevel enum.Logger, msg string) {

	if logLevel != enum.INFO {
		i.next.Log(logLevel, msg)
		return
	}
	fmt.Println("INFO :: ", msg)
}

func (i *Info) setNext(l Logger) {
	i.next = l
}
