package main

import (
	"github.com/gotutorial/oop-with-go/encapsulation/interfaces/payment"
)

func main() {

	var option payment.PaymentOption

	option = payment.CreateCreditAccount(
		"John Smith",
		"1111-2222-3333-4444",
		4,
		2023,
		123)

	option.ProcessPayment(340)

	option = payment.CreateCashAccount()

	option.ProcessPayment(120)
}
