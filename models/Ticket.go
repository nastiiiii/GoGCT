package models

// @TODO  figure out enum for tickets
// @TODO  json back and forth

type BookingStatus string

const (
	Booked BookingStatus = "Booked"
	Payed  BookingStatus = "Payed"
	Taken  BookingStatus = "Taken"
)

type Ticket struct {
	TicketId      int
	TransactionId int
	Seat          string
	PerformanceId int
	TicketStatus  BookingStatus
}

type ITicketService interface {
	getPrice(performance Performance, performanceSeat string) float64
	getPriceByTicket(ticket Ticket) float64
	changeSeatAvailability(seat string, status BookingStatus)
}
