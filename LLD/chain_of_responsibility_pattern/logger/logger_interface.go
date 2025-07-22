package logger

import "solid_design/chain_of_responsibility/enum"

type Logger interface {
	Log(logLevel enum.Logger, msg string)
	setNext(l Logger)
}


func GetLogger()  Logger {
	err := &Error{}
	debug := &Debug{}
	info := &Info{}
	info.setNext(debug)
	debug.setNext(err)
	return info
}