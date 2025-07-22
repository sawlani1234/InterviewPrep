package paymentstrategy

import "solid_design/parking_lot/entity"



func NewPaymentStrategy(paymentType string) entity.PaymentStrategy {
	switch paymentType {
	case "credit_card":
		return CreditCard{}
	case "cash":
		return Cash{}
	case "upi":
		return UPI{}
	default:
		return nil
	}
}
