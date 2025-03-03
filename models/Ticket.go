package models

// @TODO  figure out enum for tickets
// @TODO  json back and forth

type BookingStatus int

const (
	Booked BookingStatus = iota
	Payed
	Taken
)

type Ticket struct {
	//ticketId      int
	transactionId int
	seat          string
	performanceId int
	bookingStatus BookingStatus
}

type ITicketService interface {
	getPrice(performance Performance, performanceSeat string) float64
	getPriceByTicket(ticket Ticket) float64
	changeSeatAvailability(seat string, status BookingStatus)
}
