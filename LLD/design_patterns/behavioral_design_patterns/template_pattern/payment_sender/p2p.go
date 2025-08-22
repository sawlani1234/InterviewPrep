package paymentsender

import "fmt"


type P2P struct {
}

func NewP2P() P2P {
	return P2P{}
}


func (p P2P) validateRequest() {
	fmt.Println("Inside P2P method")
}

func (p P2P) debitMoney() {
	fmt.Println("debiting money from p2p method")
}

func(p P2P) deductFee() {
	fmt.Println("deducting fee from p2p method")
}


func (p P2P) creditMoney() {
	fmt.Println("crediting money using p2p method")
}