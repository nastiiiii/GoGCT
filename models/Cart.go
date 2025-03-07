package models

import "github.com/jackc/pgx/v5"

type CartService struct {
	dB       *pgx.Conn
	userCart map[int][]Ticket
}

type ICartService interface {
	AddTickets([]Ticket)
	RemoveTicket(id int)
	GetTickets() []Ticket
	ClearCart()
	CalculateTotalPrice() float64
}

func (cs *CartService) AddTickets(tickets []Ticket) {
	//TODO implement me
	panic("implement me")
}

func (cs *CartService) RemoveTicket(id int) {
	//TODO implement me
	panic("implement me")
}

func (cs *CartService) GetTickets() []Ticket {
	//TODO implement me
	panic("implement me")
}

func (cs *CartService) ClearCart() {
	//TODO implement me
	panic("implement me")
}

func (cs *CartService) CalculateTotalPrice() float64 {
	//TODO implement me
	panic("implement me")
}
