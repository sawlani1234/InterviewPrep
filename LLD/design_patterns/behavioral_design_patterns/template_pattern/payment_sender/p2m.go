package paymentsender
import "fmt"

type P2M struct {
}

func NewP2M() P2M {
	return P2M{}
}

func (p P2M) validateRequest() {
	fmt.Println("Inside P2M method")
}

func (p P2M) debitMoney() {
	fmt.Println("debiting money from P2M method")
}

func(p P2M) deductFee() {
	fmt.Println("deducting fee from P2M method")
}


func (p P2M) creditMoney() {
	fmt.Println("crediting money using P2M method")
}