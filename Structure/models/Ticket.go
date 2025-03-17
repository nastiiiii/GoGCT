package models

type BookingStatus string

const (
	Booked BookingStatus = "Booked"
	Free   BookingStatus = "Free"
	Payed  BookingStatus = "Payed"
)

type Ticket struct {
	TicketId      int
	TransactionId int
	PerformanceId int
	TicketStatus  BookingStatus
	Seat          string
}
