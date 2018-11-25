package main

import (
	"fmt"

	"github.com/gotutorial/oop-with-go/encapsulation/pod/payment"
)

func main() {

	credit := payment.CreateCreditAccount(
		"John Smith",
		"1111-2222-3333-4444",
		4,
		2023,
		123)

	fmt.Printf("Card Owner Name: %v\n", credit.OwnerName())
	fmt.Printf("Card Number: %v\n", credit.CardNumber())
	fmt.Printf("\n")
	newCardNumber := "1234-5678-01"
	fmt.Printf("Trying to change the card number to %v ...\n", newCardNumber)
	err := credit.SetCardNumber(newCardNumber)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Printf("\n")
	fmt.Printf("Available Credit: %v\n", credit.AvailableCredit())
}
