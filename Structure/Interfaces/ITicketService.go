package Interfaces

import Models "GCT/Structure/models"

// Do not forget to chnage aviability

// ITicketService Description: Shows all methods which TicketService is implemented and to which the controllers have access
type ITicketService interface {
	//Create
	CreateTicket(ticket Models.Ticket) (*Models.Ticket, error)
	//GetPrice
	GetPriceByTicket(ticket Models.Ticket) float64
	GetTicketsPriceByTransaction(transactionId int) float64
	//GetTicket
	GetTicketById(ticketId int) (*Models.Ticket, error)
	GetTicketsByTransactionId(transactionId int) ([]Models.Ticket, error) //By transaction
	GetTicketsByAccount(token string) ([]Models.Ticket, error)            //By Account
	//Delete
	DeleteTicket(ticketId int) bool
	DeleteTicketsByTransactionId(transactionId int) bool
	//Update
	UpdateTicket(id int, ticket Models.Ticket) (*Models.Ticket, error)
}
