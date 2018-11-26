package main

import "fmt"

type Account struct{}

func (a *Account) AvailableFunds() {
	fmt.Println("Listing available funds")
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Printf("Processing payment for $%v \n", amount)
	return true
}

type CreditAccount struct {
	Account
}

func main() {
	ca := &CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(300)
}
