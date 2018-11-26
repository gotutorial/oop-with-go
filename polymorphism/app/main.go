package main

import (
	"github.com/gotutorial/oop-with-go/polymorphism/payment"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

func main() {

	var option PaymentOption

	option = payment.CreateCreditAccount(
		"John Smith",
		"1111-2222-3333-4444",
		4,
		2023,
		123)

	option.ProcessPayment(340)

	option = payment.CreateCashAccount()

	option.ProcessPayment(120)

	option = payment.CreateDebitCard(
		"1234-1234-1234-1234",
		"9876",
	)

	option.ProcessPayment(50)
}
