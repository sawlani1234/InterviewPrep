package main 

import "bridgepattern/connector"

/*
1. So use bridge pattern when you have orthogonal type of classes which can grow 
independently

2. Taking below ex we have different type of computers and can connect to any printer
if we creaate all type of classes under computer interface there will be explosition 
of classes n*m*...

3. so here bridge pattern says to distribute this independent classes into differnt abstraction (interfaces)
and like printer into interfaces and computer only interacts with this interface as a conrtacct

4. now any class can be linked to any other implementation thus helping in preventing
explosion

*/

func main() {
	connector.NewMac(connector.Epson{}).PrintDoc()
	connector.NewMac(connector.Hp{}).PrintDoc()
	connector.NewWindows(connector.Hp{}).PrintDoc()
}