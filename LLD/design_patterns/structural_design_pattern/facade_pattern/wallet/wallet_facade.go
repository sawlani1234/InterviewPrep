package wallet

import "fmt"


type WalletFacade struct {
	account *Account
	securityCheck *SecurityCheck
	wallet *Wallet 
}


func NewWalletFacade(account *Account,securityCheck *SecurityCheck,wallet *Wallet) *WalletFacade {
	return &WalletFacade{
		account,
		securityCheck,
		wallet,
	}
}

func (w WalletFacade) validate(code int,acc string) bool {
	if !w.account.validateAccount(acc) {
		fmt.Println("Invalid Account ",acc) 
		return  false
	}

	if !w.securityCheck.checkCode(code) {
		fmt.Println("invalid code", code)
		return false
	}

	return true

}

func (w WalletFacade) AddMoneytoWalet(amount int,code int,acc string)  {

	if !w.validate(code,acc) {
		return 
	}

	w.wallet.creditBalance(amount)
	fmt.Println("Added ",amount, "  to wallet for customer ,",acc)
	
}

func (w WalletFacade) DebitMoneyFromWallet(amount int,code int,acc string)  {

	if !w.validate(code,acc) {
		return 
	}

	w.wallet.debitBalance(amount)
	
}