package payment

import (
	"fmt"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type CashAccount struct{}

func (c *CashAccount) ProcessPayment(amount float32) bool {
	fmt.Printf("Processing a cash transaction ...\n")
	fmt.Printf("$%v Submitted\n", amount)
	return true
}

type CreditCard struct {
	OwnerName       string
	CardNumber      string
	ExpirationMonth int
	ExpirationYear  int
	SecurityCode    int
}

func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Printf("Processing a credit card transaction ...\n")
	fmt.Printf("$%v Submitted\n", amount)
	return true
}
