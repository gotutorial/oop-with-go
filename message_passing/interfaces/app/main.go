package main

import (
	"github.com/gotutorial/oop-with-go/message_passing/interfaces/payment"
)

func main() {

	var option payment.PaymentOption

	option = &payment.CreditCard{
		OwnerName:       "John Smith",
		CardNumber:      "1111-2222-3333-4444",
		ExpirationMonth: 4,
		ExpirationYear:  2023,
		SecurityCode:    123}

	option.ProcessPayment(340)

	option = &payment.CashAccount{}

	option.ProcessPayment(120)
}
