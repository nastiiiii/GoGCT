package main

import (
	Controllers "GCT/Structure/Controller"
	"GCT/Structure/Services"
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

	router := gin.Default()
	//router.GET("/users", getUsers)

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
	//todo interface change Service

	accountService := Services.AccountService{DB: conn}
	Controllers.SetupAccountRouter(router, accountService)

	reviewService := Services.ReviewService{DB: conn}
	Controllers.SetupReviewRouters(router, reviewService)

	shipmentService := Services.ShipmentService{DB: conn}
	Controllers.SetupShipmentRoutes(router, shipmentService)

	performanceService := Services.PerformanceService{DB: conn}
	Controllers.SetUpPerformanceRouters(router, performanceService)

	ticketService := Services.TicketService{DB: conn}
	Controllers.SetupTicketRoutes(router, ticketService)

	transactionService := Services.TransactionService{DB: conn}
	Controllers.SetupTransactionRoutes(router, &transactionService)

	// Start the server
	log.Println("Server is running on port 8080...")
	router.Run("localhost:8000")
}
