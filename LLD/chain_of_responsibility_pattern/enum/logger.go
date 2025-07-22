package enum

type Logger int

const (
	INVALID Logger = iota 

	INFO Logger = 1

	DEBUG Logger = 2

	ERROR Logger = 3

)