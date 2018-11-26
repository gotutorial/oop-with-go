package payment

import (
	"fmt"
)

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
}

func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

func (c CreditCard) ProcessPayment(amount float32) bool {
	fmt.Printf("Processing a credit card transaction ...\n")
	fmt.Printf("$%v Submitted to %v Debit Card\n", amount, c.cardNumber)
	fmt.Println("----------------------------------")
	return true
}
