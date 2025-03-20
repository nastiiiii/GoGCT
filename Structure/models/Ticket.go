package models

type BookingStatus string

// Represents ticket statuses
const (
	Booked BookingStatus = "Booked"
	Payed  BookingStatus = "Payed"
)

// Ticket represents entity from the database
type Ticket struct {
	TicketId      int
	TransactionId int
	PerformanceId int
	TicketStatus  BookingStatus
	Seat          string
}

func NewTicket(transactionId int, performanceId int, seat string) Ticket {
	return Ticket{
		TransactionId: transactionId,
		PerformanceId: performanceId,
		TicketStatus:  Booked,
		Seat:          seat,
	}
}
