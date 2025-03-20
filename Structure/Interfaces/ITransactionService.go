package Interfaces

import (
	"GCT/Structure/Services"
	Models "GCT/Structure/models"
)

// ITransactionService Description: Shows all methods which TransactionService is implemented and to which the controllers have access
type ITransactionService interface {
	//Create
	CreateTransaction(transaction Models.Transaction) (int, error)
	//Update
	UpdateTransaction(transactionId int, transaction Models.Transaction) (Models.Transaction, error)
	//Delete
	DeleteTransaction(transactionId int) bool
	//Get
	GetTransactionById(transactionId int) (Models.Transaction, error)
	GetTransactionByAccount(accountId int) ([]Models.Transaction, error)
	GetTransactionsByStatus(transactionStatus Models.TransactionStatus) []Models.Transaction
	GetByConfirmationId(confirmation string) (Models.Transaction, error)
	//Payment related functions
	ProcessTransactionPayment(transactionId int) error
	SetPaymentMethod(strategy Services.PaymentStrategy)
	//Discount related functions
	//TODO Apply Discounts
}
