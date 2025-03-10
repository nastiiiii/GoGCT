package Services

import (
	Models "GCT/Structure/models"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var jwtToken = []byte("super-secret-token-for-testing-jwt@@@")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{Username: username, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtToken)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

type AccountService struct {
	DB *pgx.Conn
}

func (a *AccountService) Register(account Models.Account) (Models.Account, error) {
	query := `INSERT INTO account (contact_info, is_social_club, user_dob, username, account_hashed_password, account_balance)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING account_id`

	err := a.DB.QueryRow(
		context.Background(),
		query,
		account.ContactInfo,
		account.IsSocialClub,
		account.UserDOB,
		account.Username,
		account.AccountHashedPassword,
		account.AccountBalance,
	).Scan(&account.AccountId)

	if err != nil {
		return Models.Account{}, errors.New("failed to register account")
	}
	return account, nil
}

func (a *AccountService) CreateAccountByParams(contactInfo string, isSocialClub bool, userDOB time.Time, username string, password string) (Models.Account, error) {
	account := Models.NewAccount(contactInfo, isSocialClub, userDOB, username, password)
	return a.Register(account)
}

func (a *AccountService) Login(username string, password string) (string, error) {
	var account Models.Account
	query := `SELECT account_id, contact_info, is_social_club, user_dob, username, account_hashed_password, account_balance 
	          FROM account WHERE username = $1`
	err := a.DB.QueryRow(context.Background(), query, username).Scan(
		&account.AccountId,
		&account.ContactInfo,
		&account.IsSocialClub,
		&account.UserDOB,
		&account.Username,
		&account.AccountHashedPassword,
		&account.AccountBalance,
	)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.AccountHashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := GenerateToken(username)
	if err != nil {
		return "", errors.New("could not generate token")
	}
	return token, nil
}

func (a *AccountService) Logout() bool {
	// Since JWT tokens are stateless, logging out is typically done on the frontend by removing the token.
	return true
}

func (a *AccountService) GetAccountById(id int) (Models.Account, error) {
	var account Models.Account
	query := `SELECT account_id, contact_info, is_social_club, user_dob, username, account_hashed_password, account_balance
	          FROM account WHERE account_id = $1`
	err := a.DB.QueryRow(context.Background(), query, id).Scan(
		&account.AccountId,
		&account.ContactInfo,
		&account.IsSocialClub,
		&account.UserDOB,
		&account.Username,
		&account.AccountHashedPassword,
		&account.AccountBalance,
	)
	if err != nil {
		return Models.Account{}, errors.New("account not found")
	}
	return account, nil
}

func (a *AccountService) UpdateAccount(id int, account Models.Account) (Models.Account, error) {
	query := `UPDATE account SET contact_info = $1, is_social_club = $2, user_dob = $3, username = $4, account_balance = $5
	          WHERE account_id = $6 RETURNING account_id`

	err := a.DB.QueryRow(
		context.Background(),
		query,
		account.ContactInfo,
		account.IsSocialClub,
		account.UserDOB,
		account.Username,
		account.AccountBalance,
		id,
	).Scan(&account.AccountId)

	if err != nil {
		return Models.Account{}, errors.New("failed to update account")
	}
	return account, nil
}

func (a *AccountService) DeleteAccount(id int) error {
	query := `DELETE FROM account WHERE account_id = $1`
	_, err := a.DB.Exec(context.Background(), query, id)
	if err != nil {
		return errors.New("failed to delete account")
	}
	return nil
}

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

func (a *AccountService) HasAttendedThePerformance(accountId int, performanceId int) (bool, error) {
	query := `SELECT COUNT(*) FROM tickets WHERE account_id = $1 AND performance_id = $2 AND ticket_status = $3`
	var count int
	err := a.DB.QueryRow(context.Background(), query, accountId, performanceId, Models.Payed).Scan(&count)
	if err != nil {
		return false, errors.New("error checking attendance")
	}
	return count > 0, nil
}
