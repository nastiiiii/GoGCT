package models

type TransactionStatus string

const (
	Processing TransactionStatus = "Processing"
	Completed  TransactionStatus = "Completed"
	Canceled   TransactionStatus = "Canceled"
)

// Transaction struct
type Transaction struct {
	TransactionID     int
	ShipmentId        int
	AccountId         int
	TransactionStatus TransactionStatus
	ConfirmationID    string
	TotalCost         float64
}
