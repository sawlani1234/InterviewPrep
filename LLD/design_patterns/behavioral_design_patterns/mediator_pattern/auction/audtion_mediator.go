package auction 


type AuctionMediator interface {
	addBidder(c Collegue)
    addBid(c Collegue,amount int) 
}

