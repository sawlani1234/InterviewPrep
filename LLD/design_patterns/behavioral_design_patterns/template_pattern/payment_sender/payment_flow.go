package paymentsender


type PaymentFlow interface {
	validateRequest()
	debitMoney() 
	deductFee()
	creditMoney()
}




