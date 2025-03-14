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

	var shipService Structure.IShipmentService
	shipService = Structure.ShipmentService{DB: conn}

	shipment := Models.NewShipment(time.Now(), "address", Models.BookingStatus("Payed"), false)

	newShipment, err := shipService.CreateShipment(*shipment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newShipment)

	getShipment, err := shipService.GetShipmentById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getShipment)

	shipmentUpdate := Models.NewShipment(time.Now(), "NEW", Models.BookingStatus("Payed"), true)

	updateShipment, err := shipService.UpdateShipment(*shipmentUpdate, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateShipment)

	getShipment, err = shipService.GetShipmentById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getShipment)

	deleteShipment := shipService.DeleteShipment(1)
	fmt.Println(deleteShipment)

}
