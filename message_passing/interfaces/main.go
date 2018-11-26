package main

import "fmt"

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

func main() {

	var option PaymentOption

	option = &CreditCard{
		OwnerName:       "John Smith",
		CardNumber:      "1111-2222-3333-4444",
		ExpirationMonth: 4,
		ExpirationYear:  2023,
		SecurityCode:    123}

	option.ProcessPayment(340)

	option = &CashAccount{}

	option.ProcessPayment(120)
}
