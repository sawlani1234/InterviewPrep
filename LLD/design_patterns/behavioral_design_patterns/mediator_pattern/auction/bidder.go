package auction

import "fmt"

type Bidder struct {
	id              int
	name            string
	auctionMediator AuctionMediator
}

func NewBidder(id int, name string, auctionMediator AuctionMediator) Bidder {
	obj := Bidder{id, name, auctionMediator}
	auctionMediator.addBidder(obj)
	return obj
}

func (b Bidder) AddBid(amount int) {
	b.auctionMediator.addBid(b, amount)
}

func (b Bidder) getId() int {
	return b.id
}

func (b Bidder) recieveNotifcation(amount int) {
	fmt.Println("Current bid amount notified to bidder ", b.name, " is ", amount)
}
