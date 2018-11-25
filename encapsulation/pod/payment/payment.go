package payment

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	owenerName      string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
}

func CreateCreditAccount(owenerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		owenerName:      owenerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

func (c CreditCard) OwnerName() string {
	return c.owenerName
}

func (c CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid owener name provided")
	}
	c.owenerName = value
	return nil
}

func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

func (c *CreditCard) SetCardNumber(value string) error {
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid credit card number format")
	}
	c.cardNumber = value
	return nil
}

func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

func (c *CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() ||
		year == now.Year() && time.Month(month) < now.Month() {
		return errors.New("Expiration Date must be in the future")
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

func (c *CreditCard) SetCreditCard(value int) error {
	if value < 100 || value > 999 {
		return errors.New("Security code must be a three digit number between 100-999")
	}
	c.securityCode = value
	return nil
}

func (c CreditCard) AvailableCredit() float32 {
	return 5000
}
