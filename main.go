package main

import (
	Structure "GCT/Structure/Services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

// all the models requires the getters and setters and basic constructor
type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []user{
	{Id: 546, Username: "John"},
	{Id: 894, Username: "Mary"},
	{Id: 326, Username: "Jane"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {

	//router := gin.Default()
	//router.GET("/users", getUsers)
	//router.Run("localhost:8000")

	connStr := "postgres://root:beetroot@localhost:5433/GCT"

	var err error
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var testValue string
	err = conn.QueryRow(context.Background(), "select  'Connection string successful'").Scan(&testValue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(testValue)

	var account Structure.AccountService
	account.DB = conn
	//newAccount := Model.NewAccount("contact", false, time.Now(), "username", "password")

	/*acc, err := account.Register(newAccount)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acc)

	stillAccount, err := account.GetAccountById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stillAccount)*/

	/*token, err := account.Login("username", "password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)

	newAccount, err := account.GetUserByToken(token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newAccount)

	updatedAccount := models.NewAccount("NEW CONTACT", true, time.Now(), "New user", "password")
	fmt.Println(updatedAccount)

	newUpdatedAccount, err := account.UpdateAccount(1, updatedAccount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUpdatedAccount)
	
	err = account.DeleteAccount(1)
	if err != nil {
		log.Fatal(err)
	}
	*/
	stillAccount, err := account.GetAccountById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stillAccount)

}
