package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

//all the models requires the getters and setters and basic constructor

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

}
