package models

import (
	"github.com/jackc/pgx/v5"
	"time"
)

type Account struct {
	accountId             int
	contactInfo           string
	isSocialClub          bool
	userDOB               time.Time
	username              string
	accountBalance        float64
	accountHashedPassword string
	service               IAccountService
}

type AccountService struct {
	DB *pgx.Conn
}

type IAccountService interface {
	CreateAccount(account Account) Account
	CreatAccountByParams(contactInfo string, isSocialClub bool, userDOB time.Time, username string, password string)
	Login(username string, password string) bool
	Logout() bool
	getAccountById(id int) Account
	getTickets() []Ticket
	hasAttendedThePerformance(performance Performance) bool
	UpdateAccount(id int, account Account) Account
	DeleteAccount(id int) Account
}

func (a *AccountService) CreateAccount(account Account) Account {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) Login(username string, password string) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) Logout() bool {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) getAccountById(id int) Account {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) hasAttendedThePerformance(performance Performance) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) UpdateAccount(id int, account Account) Account {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) DeleteAccount(id int) Account {
	//TODO implement me
	panic("implement me")
}
