package main

import (
	"GCT/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

func main() {
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
	newShipment := models.Shipment{
		DateOfDispatch:  time.Now(),
		ShippingAddress: "1234 Elm Street, New York, NY",
		ShipmentStatus:  models.Pending, // Using enum-like constants
		IsUrgent:        true}

	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO "Shipments" (date_of_dispatch, shipping_address, shipment_status, is_urgent) VALUES ($1, $2, $3, $4)`,
		newShipment.DateOfDispatch,
		newShipment.ShippingAddress,
		newShipment.ShipmentStatus,
		newShipment.IsUrgent,
	)
	if err != nil {
		log.Fatal("Insert failed:", err)
	}

	fmt.Println("Shipment inserted successfully!")
}
