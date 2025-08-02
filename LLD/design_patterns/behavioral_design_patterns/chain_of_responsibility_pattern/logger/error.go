package logger

import (
	"fmt"
	"solid_design/chain_of_responsibility/enum"
)

type Error struct {
	next Logger
}

func (e *Error) Log(logLevel enum.Logger, msg string) {

	if logLevel != enum.ERROR {
		fmt.Println("Invalid logger requested")
	}
	fmt.Println("Error :: ", msg)
}

func (e *Error) setNext(l Logger) {
	e.next = l
}
