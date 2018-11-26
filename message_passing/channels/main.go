package main

import (
	"fmt"
	"strconv"
)

type CreditAccount struct{}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Printf("\nProcessing a credit card transaction ...\n")
	fmt.Printf("$%v sumbitted \n", amount)
	fmt.Printf("Enter the payment amount: ")
}

func CreateCreditAccount(chargeChannel chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeChannel)

	return creditAccount
}

func main() {
	chargeChannel := make(chan float32)
	CreateCreditAccount(chargeChannel)
	fmt.Printf("Enter the payment amount: ")
	var a string
	for a != "q" {
		fmt.Scanln(&a)
		value, err := strconv.ParseFloat(a, 32)
		if err == nil {
			chargeChannel <- float32(value)
		}
	}
}
