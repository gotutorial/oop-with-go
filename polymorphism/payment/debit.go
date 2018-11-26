package payment

import (
	"fmt"
)

type DebitCard struct {
	cardNumber string
	pin        string
}

func CreateDebitCard(cardNumber string, pin string) *DebitCard {
	return &DebitCard{
		cardNumber: cardNumber,
		pin:        pin,
	}
}

func (d DebitCard) ProcessPayment(amount float32) bool {
	fmt.Printf("Processing a debit card transaction ...\n")
	fmt.Printf("$%v Submitted to %v Debit Card\n", amount, d.cardNumber)
	fmt.Println("----------------------------------")
	return true
}
