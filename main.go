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

	var perService Structure.IPerformanceService
	perService = Structure.PerformanceService{DB: conn}

	//performance := Models.NewPerformance("A:100; B:200; C:300;", "A: 12-true, 13-false, 14-true; B: 15-true, 16-true, 17-true; C: 18-false, 19-false, 20-true;", "NEW", "Descriptive", "actor 2", time.Now())

	/*newPerformance, err := perService.CreatePerformance(performance)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newPerformance)

	getPerformance, err := perService.GetPerformanceById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getPerformance)*/
	/*updatePerformance, err := perService.UpdatePerformance(performance, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updatePerformance)

	getByName, err := perService.GetPerformanceByName("NEW")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getByName)
	*/
	//deletedPerformance := perService.DeletePerformance(5)
	//fmt.Println(deletedPerformance)

	getPerformance, err := perService.GetPerformanceById(1)
	if err != nil {
		log.Fatal(err)
	}

	getAllSeats, err := perService.GetAllSeats(getPerformance)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getAllSeats)

	getAllAvailableSeats, err := perService.GetAvailableSeats(getPerformance)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getAllAvailableSeats)

	getSeatPriceB, err := perService.GetSeatPrice(getPerformance, "B")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getSeatPriceB)

	getPerformancePrice, err := perService.GetPerformancePrice(getPerformance)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getPerformancePrice)

	changeA12Seat, err := perService.ChangeSeatAvailability(&getPerformance, "A", 14, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(changeA12Seat)

	getAllSeats, err = perService.GetAllSeats(getPerformance)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getAllSeats)
}
