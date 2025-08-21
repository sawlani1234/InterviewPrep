package main 

import "facade_pattern/wallet"

/*
Facade pattern is basically used to hide complexities of subssytem classes , 
it basical wraps all the complex steps of subsyem classes and provides simple methods
to interact with subsytem

for eg in below eg it hides the complex checks of add money to wallet checks and a provide 
simple functiont to client to add money with the paramteres 
*/

func main() {
		walletFacade := wallet.NewWalletFacade(
			wallet.NewAccount("shubham"),
			wallet.NewSecurityCheck(120),
			wallet.NewWallet(),
		)

		walletFacade.AddMoneytoWalet(12,100,"nishi")
		walletFacade.AddMoneytoWalet(10,120,"shubham")
}