package main

import (
	Structure "GCT/Structure/Services"
	Models "GCT/Structure/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"time"
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

	accountService := Structure.AccountService{DB: conn}
	account := Models.NewAccount("contact", false, time.Now(), "username", "pass")

	newAccount, err := accountService.Register(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newAccount)

	var reviewService Structure.IReviewService
	reviewService = Structure.ReviewService{DB: conn}

	review := Models.NewReview(2, 2, "review", 4, time.Now())

	createReview, err := reviewService.CreateReview(*review)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(createReview)

	getReviewByAccount, err := reviewService.GetReviewsByAccountId(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getReviewByAccount)

	getReviewByPerformance, err := reviewService.GetReviewsByPerformanceId(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getReviewByPerformance)

	err = reviewService.DeleteReview(5)
	if err != nil {
		log.Fatal(err)
	}

	//192.168.108.54

}
