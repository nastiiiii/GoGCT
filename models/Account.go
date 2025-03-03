package models

import (
	"time"
)

type Account struct {
	// accountId int
	contactInfo           string
	isSocialClub          bool
	userDOB               time.Time
	username              string
	accountBalance        float64
	accountHashedPassword string
}

type IAccountService interface {
	registerAccount(account Account) Account
	login(username string, password string)
	logout()
	getAccountById(id int) Account
	getTickets() []Ticket
	hasAttendedThePerformance(performance Performance) bool
}
