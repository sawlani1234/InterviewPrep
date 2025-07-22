package main

import (
	"fmt"
	"solid_design/chain_of_responsibility/enum"
	"solid_design/chain_of_responsibility/logger"
)


func main() {
	
	fmt.Println("here")
	
	logger := logger.GetLogger()
	logger.Log(enum.INFO,"INFO LOG")
	logger.Log(enum.DEBUG,"DEBUG LOG")
	logger.Log(enum.ERROR,"error LOG")
}



// func message(log level)