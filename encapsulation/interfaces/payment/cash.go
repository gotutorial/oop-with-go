package payment

import "fmt"

type Cash struct{}

func CreateCashAccount() *Cash {
	return &Cash{}
}

func (c Cash) ProcessPayment(amount float32) bool {
	fmt.Printf("Processing a cash transaction ...\n")
	fmt.Printf("$%v Submitted\n", amount)
	return true
}
