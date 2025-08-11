package main

import "fmt"
import "adapter_design_pattern/weight"

// Wright example

/*
1. Adapter pattern is used when we need conversion between objects
2. Famously lets say we integrate to a 3rd party which uses JSON as communication
and our code is primary XML based then a converter interface acting as a bridge
can help here

3. The realest world analogy is the charger adapter used for connecting to socket



*/
func main() {

	fmt.Println(weight.NewKgAdapter(weight.NewLbsImpl(23)).GetWeightInKg())
   

}