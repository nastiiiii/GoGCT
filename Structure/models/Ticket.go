package models

type BookingStatus string

const (
	Booked BookingStatus = "Booked"
	Payed  BookingStatus = "Payed"
)

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
