package logger

import (
	"fmt"
	"solid_design/chain_of_responsibility/enum"
)

type Debug struct {
	next Logger
}

func (d *Debug) Log(logLevel enum.Logger, msg string) {

	if logLevel != enum.DEBUG {
		d.next.Log(logLevel, msg)
		return
	}
	fmt.Println("Debug :: ", msg)
}

func (d *Debug) setNext(l Logger) {
	d.next = l
}
