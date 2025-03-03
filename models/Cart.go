package models

type Cart struct {
	userCart map[int][]Ticket
}

type ICartService interface {
	addTickets([]Ticket)
	removeTicket(id int)
	getTickets() []Ticket
	clearCart()
	calculateTotalPrice() float64
}
