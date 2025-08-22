package paymentsender


type PaymmentTemplate struct {
	paymentFlow PaymentFlow
}

func NewPaymentTemplate(paymentFlow PaymentFlow) PaymmentTemplate {
	return PaymmentTemplate{paymentFlow}
}


func (p PaymmentTemplate) ProcessPayment() {
	p.paymentFlow.validateRequest()
	p.paymentFlow.debitMoney()
	p.paymentFlow.deductFee()
	p.paymentFlow.creditMoney()
}