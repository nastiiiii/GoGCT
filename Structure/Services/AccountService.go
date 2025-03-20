package Services

import (
	Middleware "GCT/Structure/middleware"
	Models "GCT/Structure/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// AccountService implements the database operations and businesses logic related to Account
type AccountService struct {
	DB *pgx.Conn
}

func (a *AccountService) Register(account Models.Account, password string) (Models.Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	account.AccountHashedPassword = string(hashedPassword)
	query := `INSERT INTO "Accounts" ("contactInfo", "isSocialClub", "userDOB", username, "accountBalance", "accountHashedPassword")
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING "accountID"`

	err = a.DB.QueryRow(
		context.Background(),
		query,
		account.ContactInfo,
		account.IsSocialClub,
		account.UserDOB,
		account.Username,
		account.AccountBalance,
		account.AccountHashedPassword,
	).Scan(&account.AccountId)

	if err != nil {
		return Models.Account{}, errors.New("failed to register account")
	}
	return account, nil
}

func (a *AccountService) CreateAccountByParams(contactInfo string, isSocialClub bool, userDOB time.Time, username string, password string) (Models.Account, error) {
	account := Models.NewAccount(contactInfo, isSocialClub, userDOB, username, password)
	return a.Register(account, password)
}

func (a *AccountService) Login(username string, password string) (string, error) {
	var account Models.Account
	query := `SELECT * FROM "Accounts" WHERE username = $1`
	err := a.DB.QueryRow(context.Background(), query, username).Scan(
		&account.AccountId,
		&account.IsSocialClub,
		&account.UserDOB,
		&account.Username,
		&account.AccountBalance,
		&account.AccountHashedPassword,
		&account.ContactInfo,
	)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.AccountHashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := Middleware.GenerateToken(username)
	if err != nil {
		return "", errors.New("could not generate token")
	}
	return token, nil
}

func (a *AccountService) GetUserByToken(tokenString string) (Models.Account, error) {
	claims, err := Middleware.ParseToken(tokenString)
	if err != nil {
		return Models.Account{}, errors.New("unauthorized: invalid token")
	}

	var account Models.Account
	query := `SELECT * FROM "Accounts" WHERE username = $1`
	err = a.DB.QueryRow(context.Background(), query, claims.Username).Scan(
		&account.AccountId,
		&account.IsSocialClub,
		&account.UserDOB,
		&account.Username,
		&account.AccountBalance,
		&account.AccountHashedPassword,
		&account.ContactInfo,
	)

	if err != nil {
		return Models.Account{}, errors.New("account not found")
	}

	return account, nil
}

func (a *AccountService) Logout() bool {
	// Since JWT tokens are stateless, logging out is typically done on the frontend by removing the token.
	return true
}

func (a *AccountService) GetAccountById(id int) (Models.Account, error) {
	var account Models.Account
	query := `SELECT * FROM "Accounts" WHERE "accountID" = $1`
	err := a.DB.QueryRow(context.Background(), query, id).Scan(
		&account.AccountId,
		&account.IsSocialClub,
		&account.UserDOB,
		&account.Username,
		&account.AccountBalance,
		&account.AccountHashedPassword,
		&account.ContactInfo,
	)
	if err != nil {
		return Models.Account{}, errors.New("account not found")
	}
	return account, nil
}

func (a *AccountService) UpdateAccount(id int, account Models.Account) (Models.Account, error) {
	query := `UPDATE "Accounts" SET "contactInfo" = $1, "isSocialClub" = $2, "userDOB" = $3, username = $4, "accountBalance" = $5
	          WHERE "accountID" = $6 RETURNING "accountID"`

	err := a.DB.QueryRow(
		context.Background(),
		query,
		&account.ContactInfo,
		&account.IsSocialClub,
		&account.UserDOB,
		&account.Username,
		&account.AccountBalance,
		id,
	).Scan(&account.AccountId)

	if err != nil {
		return Models.Account{}, errors.New("failed to update account")
	}
	return account, nil
}

func (a *AccountService) DeleteAccount(id int) error {
	query := `DELETE FROM "Accounts" WHERE "accountID" = $1`
	_, err := a.DB.Exec(context.Background(), query, id)
	if err != nil {
		return errors.New("failed to delete account")
	}
	return nil
}

// GetTickets Description: Get Tickets related exactly to the logged in account
func (a *AccountService) GetTickets(accountId int) ([]Models.Ticket, error) {
	query := `SELECT ticket_id, transaction_id, seat, performance_id, ticket_status FROM tickets WHERE account_id = $1`
	rows, err := a.DB.Query(context.Background(), query, accountId)
	if err != nil {
		return nil, errors.New("could not retrieve tickets")
	}
	defer rows.Close()

	var tickets []Models.Ticket
	for rows.Next() {
		var ticket Models.Ticket
		var ticketStatus string // Temporary variable to store ticket_status as string

		err := rows.Scan(&ticket.TicketId, &ticket.TransactionId, &ticket.Seat, &ticket.PerformanceId, &ticketStatus)
		if err != nil {
			log.Println("Error scanning ticket:", err)
			continue
		}

		// Convert ticketStatus string to BookingStatus type
		ticket.TicketStatus = Models.BookingStatus(ticketStatus)

		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

// HasAttendedThePerformance Description: Checks if user has attended the performance
func (a *AccountService) HasAttendedThePerformance(accountId int, performanceId int) (bool, error) {
	query := `SELECT COUNT(*) FROM tickets WHERE account_id = $1 AND performance_id = $2 AND ticket_status = $3`
	var count int
	err := a.DB.QueryRow(context.Background(), query, accountId, performanceId, Models.Payed).Scan(&count)
	if err != nil {
		return false, errors.New("error checking attendance")
	}
	return count > 0, nil
}
