package models

import "fmt"

type PaymentStrategy interface {
	ProcessPayment(amount float64)
}

type CardPayment struct{}

func (c *CardPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processing cash payment of $%.2f\n", amount)
}

type CashPayment struct{}

func (c *CashPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processing cash payment of $%.2f\n", amount)
}

type AccountPayment struct{}

func (a *AccountPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processing account payment of $%.2f\n", amount)
}

// TransactionService handles transactions
type TransactionService struct {
	CartService     *CartService
	PaymentMethod   PaymentStrategy
	DiscountManager *DiscountManager
}

// ProcessPayment executes the selected payment strategy
func (t *TransactionService) ProcessPayment(amount float64) {
	if t.PaymentMethod == nil {
		fmt.Println("No payment method selected!")
		return
	}
	t.PaymentMethod.ProcessPayment(amount)
}

// GenerateConfirmationNumber generates a unique confirmation
func (t *TransactionService) GenerateConfirmationNumber() string {
	return "CONF123456" // Placeholder for a unique generator
}
