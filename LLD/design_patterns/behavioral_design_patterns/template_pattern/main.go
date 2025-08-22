package main

import (
	"fmt"
	paymentsender "template_pattern/payment_sender"
)


/*
1. So in template method we take  out a set of ordered methods(rules) and push it in a template method in a class 
2. These methods are part of interface 
3. Now each client that wants to use this template method defined in template class would have to implement these rules as they are called using this interface inside the template class
4. so through this client each client can define its own impl of this set of rules and also follow this order of rules rquired by the system
5. for eg there are multiple payment methods like PAX to merchant. (P2M) and PAX to pax (p2p) , each payment requires same set of setps 
of validation , debit money , charge and credit money to destination , 
6. In above case we do not have template method we would have to add these common steps in every class , and sometimes client might 
forget to implement a crucial setp or may have the incorrect order , thus template pattern helps here in the above case 


*/

func main() {
	
	 p2p := paymentsender.NewP2P()
	 p2M := paymentsender.NewP2M()
     
	 paymentsender.NewPaymentTemplate(p2p).ProcessPayment()
	 fmt.Println()
	 paymentsender.NewPaymentTemplate(p2M).ProcessPayment()
}