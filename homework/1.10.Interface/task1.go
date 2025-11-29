package main

import (
	"fmt"
)

type PaymentProcessor interface {
	Process(amount float64) string
}

// P.S. Подумайте, какие поля могут быть у каждой структуры
type CreditCard struct {
	CardNumber string
	Owner      string
	Balance    float64
}

func (c *CreditCard) Process(amount float64) string {
	// конструкцию берем из домашки 1.9.Struct task3
	if amount > c.Balance {
		return fmt.Sprintf("Not enough money ON CreditCard %s %s. Requested %.2f$ Available %.2f$.\n",
			c.CardNumber[len(c.CardNumber)-4:], c.Owner, amount, c.Balance)
	}
	c.Balance -= amount
	return fmt.Sprintf("CreditCard %s %s. Withdraw -%.2f$. New balance: %.2f$.\n",
		c.CardNumber[len(c.CardNumber)-4:], c.Owner, amount, c.Balance)

}

type CryptoWallet struct {
	PublicKey   string
	Sender      string
	Recipient   string
	Commission  float64
	CoinBalance float64
}

func (c *CryptoWallet) Process(amount float64) string {
  commissionAmount:= amount *(c.Commission/100)
  totalAmount := amount + commissionAmount

	if totalAmount > c.CoinBalance {
		return fmt.Sprintf("Not enough coins ON CriptoWallet %s %s. Requested %.2f$ Available %.2f$.\n",
			c.PublicKey[:8], c.Sender, totalAmount, c.CoinBalance)
	}
	c.CoinBalance -= totalAmount
	return fmt.Sprintf("CryptoWallet %s %s TO %s. Withdraw -%.2f$ Commission %.2f$. New balance: %.2f$.\n",
		c.PublicKey[:8], c.Sender, c.Recipient, amount, commissionAmount, c.CoinBalance)
}

func main() {
	paymentSystems := []PaymentProcessor{
		&CreditCard{
			CardNumber: "1234567812345678",
			Owner:      "Ivanov I",
			Balance:    1500.00,
		},
		&CryptoWallet{
			PublicKey:   "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
			Sender:      "Ivanov I",
			Recipient:   "Petrov P",
			Commission:  4.5, // то есть 4.5%
			CoinBalance: 10000.5,
		},
		&CreditCard{
			CardNumber: "9876543298765432",
			Owner:      "Petrov P",
			Balance:    800.50,
		},
	}

	amount := []float64{100.123, 23.34, 99999.56}
	for _, value := range amount {
		for index, processor := range paymentSystems {
			result := processor.Process(value)
			fmt.Printf("\n%d. %s", index+1, result)
		}
	}

}
