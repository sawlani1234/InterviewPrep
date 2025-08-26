package main

import "mediator_pattern/auction"


/*
1. Mediator pattern is basically used as a contact point of  2 or more objets , 
2, each object instead of talking to each other connects with this mediator objects
3. this mediator objects basically perfomrs and delegate  action to other objects 


*/

func main() {

	ipl := auction.NewIplAuction()
	shubham := auction.NewBidder(1,"Shubham",ipl)
	nishi := auction.NewBidder(2,"Nishi",ipl)
	nishi.AddBid(2000)
	shubham.AddBid(200000000)

}