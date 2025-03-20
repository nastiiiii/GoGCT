package Services

import (
	Models "GCT/Structure/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"math/rand"
	"time"
)

// TransactionService implements the database operations and businesses logic related to Transaction
type TransactionService struct {
	DB            *pgx.Conn
	PaymentMethod PaymentStrategy
	//TODO Discount Manager
}

func (t *TransactionService) SetPaymentMethod(method PaymentStrategy) {
	t.PaymentMethod = method
}

// GenerateConfirmationNumber generates a unique confirmation
func (t *TransactionService) GenerateConfirmationNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("CONF%d", rand.Intn(1000000)) // Generates CONF followed by a random number
}

func (t *TransactionService) CreateTransaction(transaction Models.Transaction) (int, error) {
	transaction.ConfirmationID = t.GenerateConfirmationNumber()
	query := `INSERT INTO "Transactions" 
		("shipmentID", "accountID", "transactionStatus", "ConfirmationID", "totalCost") 
		VALUES ($1, $2, $3, $4, $5) RETURNING "transactionID"`
	var id int
	err := t.DB.QueryRow(
		context.Background(), query,
		transaction.ShipmentId, transaction.AccountId, transaction.TransactionStatus, transaction.ConfirmationID, transaction.TotalCost,
	).Scan(&id)

	if err != nil {
		log.Println("Error inserting performance:", err)
		return 0, err
	}
	return id, nil
}

func (t *TransactionService) UpdateTransaction(transactionId int, transaction Models.Transaction) (Models.Transaction, error) {
	query := `UPDATE "Transactions" 
		SET "shipmentID" = $1, 
		    "accountID" = $2, 
		    "transactionStatus" = $3, 
		    "ConfirmationID" = $4, 
		    "totalCost" = $5 
		WHERE "transactionID" = $6`

	_, err := t.DB.Exec(
		context.Background(), query,
		&transaction.ShipmentId,
		&transaction.AccountId,
		&transaction.TransactionStatus,
		&transaction.ConfirmationID,
		&transaction.TotalCost,
		transactionId)

	updatedTransaction, err := t.GetTransactionById(transactionId)

	if err != nil {
		log.Println("Error updating performance:", err)
		return Models.Transaction{}, err
	}
	return updatedTransaction, nil
}

func (t *TransactionService) DeleteTransaction(transactionId int) bool {
	query := `DELETE FROM "Transactions" WHERE "transactionID" = $1`
	_, err := t.DB.Exec(context.Background(), query, transactionId)
	if err != nil {
		log.Println("Error deleting transaction:", err)
	}
	return err == nil
}

func (t *TransactionService) GetTransactionById(transactionId int) (Models.Transaction, error) {
	var transaction Models.Transaction
	query := `SELECT "shipmentID", "accountID", "transactionStatus", "ConfirmationID", "totalCost" FROM "Transactions" WHERE "transactionID" = $1`

	err := t.DB.QueryRow(context.Background(), query, transactionId).Scan(
		&transaction.ShipmentId,
		&transaction.AccountId,
		&transaction.TransactionStatus,
		&transaction.ConfirmationID,
		&transaction.TotalCost,
	)
	transaction.TransactionID = transactionId
	if err != nil {
		log.Println("Error retrieving transaction:", err)
		return Models.Transaction{}, err // Return empty struct in case of failure
	}

	return transaction, nil
}

// GetTransactionsByStatus Description: to see which transactions haven't been paid yet
func (t *TransactionService) GetTransactionsByStatus(transactionStatus Models.TransactionStatus) []Models.Transaction {
	var transactions []Models.Transaction
	query := `SELECT "transactionID", "shipmentID", "accountID", "transactionStatus", "ConfirmationID", "totalCost" FROM "Transactions" WHERE "transactionStatus" = $1`

	rows, err := t.DB.Query(context.Background(), query, transactionStatus)
	if err != nil {
		log.Println("Error retrieving completed transactions:", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var model Models.Transaction
		err := rows.Scan(
			&model.TransactionID,
			&model.ShipmentId,
			&model.AccountId,
			&model.TransactionStatus,
			&model.ConfirmationID,
			&model.TotalCost)
		if err != nil {
			log.Println("Error retrieving completed transactions:", err)
			continue
		}
		transactions = append(transactions, model)
	}
	if len(transactions) == 0 {
		log.Println("No completed transactions found.")
	}

	return transactions
}

func (t *TransactionService) GetByConfirmationId(confirmationId string) (Models.Transaction, error) {
	var transaction Models.Transaction
	query := `SELECT "transactionID", "shipmentID", "accountID", "transactionStatus", "ConfirmationID", "totalCost" FROM "Transactions" WHERE "ConfirmationID" = $1`

	err := t.DB.QueryRow(context.Background(), query, confirmationId).Scan(
		&transaction.TransactionID,
		&transaction.ShipmentId,
		&transaction.AccountId,
		&transaction.TransactionStatus,
		&transaction.ConfirmationID,
		&transaction.TotalCost,
	)
	if err != nil {
		log.Println("Error retrieving transaction:", err)
		return Models.Transaction{}, err // Return empty struct in case of failure
	}

	return transaction, nil
}

func (t *TransactionService) GetTransactionByAccount(accountId int) ([]Models.Transaction, error) {
	var transactions []Models.Transaction
	query := `SELECT "transactionID", "shipmentID", "accountID", "transactionStatus", "ConfirmationID", "totalCost" FROM "Transactions" WHERE "accountID" = $1`

	rows, err := t.DB.Query(context.Background(), query, accountId)
	if err != nil {
		log.Println("Error retrieving completed transactions:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var model Models.Transaction
		err := rows.Scan(
			&model.TransactionID,
			&model.ShipmentId,
			&model.AccountId,
			&model.TransactionStatus,
			&model.ConfirmationID,
			&model.TotalCost)
		if err != nil {
			log.Println("Error retrieving completed transactions:", err)
			continue
		}
		transactions = append(transactions, model)
	}
	if len(transactions) == 0 {
		log.Println("No completed transactions found.")
		return nil, err
	}

	return transactions, nil
}

// ProcessTransactionPayment Description: to process payment by choosen method
func (t *TransactionService) ProcessTransactionPayment(transactionID int) error {
	ticketService := TicketService{DB: t.DB}
	totalCost := ticketService.GetTicketsPriceByTransaction(transactionID)
	if t.PaymentMethod == nil {
		return errors.New("No payment method")
	}
	t.PaymentMethod.ProcessPayment(totalCost)
	transaction, err := t.GetTransactionById(transactionID)
	if err != nil {
		return errors.New("Error retrieving transaction:")
	}
	transaction.TransactionStatus = Models.Completed
	transaction.TotalCost = totalCost
	_, err = t.UpdateTransaction(transactionID, transaction)
	if err != nil {
		return errors.New("Error updating transaction:")
	}
	return nil
}

// PaymentStrategy interface
type PaymentStrategy interface {
	ProcessPayment(amount float64)
}

// CardPayment struct
type CardPayment struct{}

func (c *CardPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processing card payment of $%.2f\n", amount)
}

// CashPayment struct
type CashPayment struct{}

func (c *CashPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processing cash payment of $%.2f\n", amount)
}

// AccountPayment struct
type AccountPayment struct{}

func (a *AccountPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processing account payment of $%.2f\n", amount)
}

// ProcessPayment executes the selected payment strategy
func (t *TransactionService) ProcessPayment(amount float64) {
	if t.PaymentMethod == nil {
		fmt.Println("No payment method selected!")
		return
	}
	t.PaymentMethod.ProcessPayment(amount)
}
