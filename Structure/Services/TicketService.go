package Services

import (
	Models "GCT/Structure/models"
	"context"
	"github.com/jackc/pgx/v5"
)

type TicketService struct {
	DB *pgx.Conn
}

func NewTicketService(db *pgx.Conn) *TicketService {
	return &TicketService{
		DB: db,
	}
}

// Do not forget to chnage aviability
type ITicketService interface {
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
	UpdateTicket(id int, ticket Models.Ticket)
}

func (t TicketService) CreateTicket(ticket Models.Ticket) (*Models.Ticket, error) {
	query := `INSERT INTO "Tickets" ("transactionID", "performanceID", "TicketStatus", "Seat") VALUES ($1, $2, $3, $4, $5, $6) RETURNING "ticketID"`
	var id int
	err := t.DB.QueryRow(context.Background(), query, ticket.TransactionId, ticket.PerformanceId, ticket.TicketStatus, ticket.Seat).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (t TicketService) GetPriceByTicket(ticket Models.Ticket) float64 {
	if len(ticket.Seat) == 0 {
		return 0
	}

	seatBand := string(ticket.Seat[0])
	performanceService := PerformanceService{DB: t.DB}
	performance, err := performanceService.GetPerformanceById(ticket.PerformanceId)
	if err != nil {
		return 0
	}

	price, err := performanceService.GetSeatPrice(performance, seatBand)
	if err != nil {
		return 0
	}
	return price
}

func (t TicketService) GetTicketsPriceByTransaction(transactionId int) float64 {
	var totalPrice float64
	tickets, err := t.GetTicketsByTransactionId(transactionId)
	if err != nil {
		return 0
	}
	for _, ticket := range tickets {
		totalPrice += t.GetPriceByTicket(ticket)
	}
	return totalPrice
}

func (t TicketService) GetTicketById(ticketId int) (*Models.Ticket, error) {
	var ticket Models.Ticket
	query := `SELECT "ticketID", "transactionID", "performanceID", "TicketStatus", "Seat" FROM "Tickets" WHERE "ticketID" = $1`
	err := t.DB.QueryRow(context.Background(), query, ticketId).Scan(
		&ticket.TicketId,
		&ticket.TransactionId,
		&ticket.PerformanceId,
		&ticket.TicketStatus,
		&ticket.Seat)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (t TicketService) GetTicketsByTransactionId(transactionId int) ([]Models.Ticket, error) {
	var tickets []Models.Ticket
	query := `SELECT "ticketID", "transactionID", "performanceID", "TicketStatus", "Seat" FROM "Tickets" WHERE transactionID = $1`

	rows, err := t.DB.Query(context.Background(), query, transactionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ticket Models.Ticket
		err := rows.Scan(
			&ticket.TicketId,
			&ticket.TransactionId,
			&ticket.PerformanceId,
			&ticket.TicketStatus,
			&ticket.Seat)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (t TicketService) GetTicketsByAccount(token string) ([]Models.Ticket, error) {
	accountService := AccountService{DB: t.DB}
	transactionService := TransactionService{DB: t.DB}
	account, err := accountService.GetUserByToken(token)
	if err != nil {
		return nil, err
	}

	transaction, err := transactionService.GetTransactionByAccount(account.AccountId)
	if err != nil {
		return nil, err
	}

	return t.GetTicketsByTransactionId(transaction[0].TransactionID)
}

func (t TicketService) DeleteTicket(ticketId int) bool {
	query := `DELETE FROM "Tickets" WHERE "ticketID" = $1`
	_, err := t.DB.Exec(context.Background(), query, ticketId)
	if err != nil {
		return false
	}
	return true
}

func (t TicketService) DeleteTicketsByTransactionId(transactionId int) bool {
	query := `DELETE FROM "Tickets" WHERE "transactionID" = $1`
	_, err := t.DB.Exec(context.Background(), query, transactionId)
	if err != nil {
		return false
	}
	return true
}

func (t TicketService) UpdateTicket(id int, ticket Models.Ticket) (*Models.Ticket, error) {
	query := `UPDATE "Tickets" 
		SET "transactionID" = $1, 
		    "performanceID" = $2, 
		    "TicketStatus" = $3, 
		    "Seat" = $4,
			WHERE "ticketID" = $5`
	_, err := t.DB.Exec(context.Background(), query, &ticket.TransactionId, &ticket.PerformanceId, &ticket.TicketStatus, &ticket.Seat, id)
	updatedTicket, err := t.GetTicketById(id)
	if err != nil {
		return nil, err
	}
	return updatedTicket, nil
}
