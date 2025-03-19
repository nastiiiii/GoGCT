package Interfaces

import (
	"GCT/Structure/Services"
	Models "GCT/Structure/models"
)

type ITransactionService interface {
	CreateTransaction(transaction Models.Transaction) (int, error)
	UpdateTransaction(transactionId int, transaction Models.Transaction) (Models.Transaction, error)
	DeleteTransaction(transactionId int) bool
	GetTransactionById(transactionId int) (Models.Transaction, error)
	GetTransactionByAccount(accountId int) ([]Models.Transaction, error)
	GetTransactionsByStatus(transactionStatus Models.TransactionStatus) []Models.Transaction
	GetByConfirmationId(confirmation string) (Models.Transaction, error)
	ProcessTransactionPayment(transactionId int) error
	SetPaymentMethod(strategy Services.PaymentStrategy)
	//TODO Apply Discounts
}
