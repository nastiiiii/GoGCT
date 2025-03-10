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

type IAccountService interface {
	Register(account Models.Account) Models.Account
	CreatAccountByParams(contactInfo string, isSocialClub bool, userDOB time.Time, username string, password string)
	Login(username string, password string) (string, error)
	Logout() bool
	GetAccountById(id int) Models.Account
	GetTickets() []Models.Ticket
	HasAttendedThePerformance(performance Models.Performance) bool
	UpdateAccount(id int, account Models.Account) Models.Account
	DeleteAccount(id int) Models.Account
}

func (a *AccountService) Register(account Models.Account) Models.Account {
	panic("implement me")
}

func (a *AccountService) Login(username string, password string) (string, error) {
	var account Models.Account
	query := `SELECT * FROM account WHERE username = $1`
	err := a.DB.QueryRow(context.Background(), query, username).Scan(&account)
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.AccountHashedPassword), []byte(password))

	token, err := GenerateToken(username)
	if err != nil {
		return "", errors.New("could not generate token")
	}
	return token, nil
}

func (a *AccountService) Logout() bool {

	panic("implement me")
}

func (a *AccountService) GetAccountById(id int) Models.Account {
	var account Models.Account
	query := `SELECT * FROM account WHERE id=$1`
	err := a.DB.QueryRow(context.Background(), query, id).Scan(&account)
	if err != nil {
		log.Println(err)
	}
	return account
}

func (a *AccountService) HasAttendedThePerformance(performance Models.Performance) bool {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) UpdateAccount(id int, account Models.Account) Models.Account {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) DeleteAccount(id int) Models.Account {
	//TODO implement me
	panic("implement me")
}

func (a *AccountService) CreatAccountByParams(contactInfo string, isSocialClub bool, userDOB time.Time, username string, password string) {

	panic("implement me")
}

func (a *AccountService) GetTickets() []Models.Ticket { panic("implement me") }
