package models

import "github.com/jackc/pgx/v5"

// @TODO  figure out enum for tickets
// @TODO  json back and forth

type BookingStatus string

const (
	Booked BookingStatus = "Booked"
	Free   BookingStatus = "Free"
	Payed  BookingStatus = "Payed"
)

type ITicketService interface {
	CreateTicket(transactionId int, seat string, performanceId int, ticketStatus BookingStatus) *Ticket
	GetPrice(performance Performance, performanceSeat string) float64
	GetPriceByTicket(ticket Ticket) float64
	GetTicketById(ticketId int) *Ticket
	GetAllTickets() []Ticket
	ChangeSeatAvailability(seat string, status BookingStatus)
	DeleteTicket(ticketId int) bool
	UpdateTicket(id int, ticket Ticket)
}

type Ticket struct {
	TicketId      int
	TransactionId int
	Seat          string
	PerformanceId int
	TicketStatus  BookingStatus
}

type TicketService struct {
	DB *pgx.Conn
}

func (t *TicketService) CreateTicket(transactionId int, seat string, performanceId int, ticketStatus BookingStatus) *Ticket {
	//TODO implement me
	panic("implement me")
}

func (t *TicketService) GetPriceByTicket(ticket Ticket) float64 {
	//TODO implement me
	panic("implement me")
}

func (t *TicketService) GetPrice(performance Performance, performanceSeat string) float64 {
	//TODO implement me
	panic("implement me")
}

func (t *TicketService) ChangeSeatAvailability(seat string, status BookingStatus) {
	//TODO implement me
	panic("implement me")
}

func (t *TicketService) GetTicketById(ticketId int) Ticket {
	//TODO implement me
	panic("implement me")
}

func (t *TicketService) GetAllTickets() []Ticket {
	//TODO implement me
	panic("implement me")
}

func (t *TicketService) DeleteTicket(ticketId int) bool {
	//TODO implement me
	panic("implement me")
}

func (t *TicketService) UpdateTicket(id int, ticket Ticket) {
	//TODO implement me
	panic("implement me")
}
