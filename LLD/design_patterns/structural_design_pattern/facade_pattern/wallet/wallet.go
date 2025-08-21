package wallet 


type Wallet struct {
	balance int 
}


func NewWallet() *Wallet {
	return &Wallet{}
}


func (w *Wallet) creditBalance(m int) {
	w.balance += m
}

func (w *Wallet) debitBalance(m int) {
	if m >= w.balance {
		w.balance-=m
	}
	
}

