package models

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Account struct {
	AccountId             int
	ContactInfo           string
	IsSocialClub          bool
	UserDOB               time.Time
	Username              string
	AccountBalance        float64
	AccountHashedPassword string
}

func NewAccount(contactInfo string, IsSocialClub bool, userDOB time.Time, username string, password string) Account {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return Account{
		ContactInfo:           contactInfo,
		IsSocialClub:          IsSocialClub,
		UserDOB:               userDOB,
		Username:              username,
		AccountBalance:        0,
		AccountHashedPassword: string(hashedPassword),
	}
}
