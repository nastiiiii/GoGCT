package Services

import (
	Models "GCT/Structure/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

type IShipmentService interface {
	CreateShipment(dateOfDispatch time.Time, shippingAddress string, shipmentStatus Models.BookingStatus, isUrgent bool) (*Models.Shipment, error)
	UpdateShipment(shipment Models.Shipment, id int) Models.Shipment
	DeleteShipment(id int)
	GetShipmentById(id int) Models.Shipment
}

type ShipmentService struct {
	DB *pgx.Conn
}

func (s ShipmentService) CreateShipment(dateOfDispatch time.Time, shippingAddress string, shipmentStatus Models.BookingStatus, isUrgent bool) (*Models.Shipment, error) {
	var err error
	newShipment := &Models.Shipment{
		DateOfDispatch:  dateOfDispatch,
		ShippingAddress: shippingAddress,
		ShipmentStatus:  shipmentStatus,
		IsUrgent:        isUrgent,
	}
	var shipmentId int
	err = s.DB.QueryRow(
		context.Background(),
		`INSERT INTO "Shipments" ("dateOfDispatch", "shippingAddress", "shipmentStatus", "isUrgent") VALUES ($1, $2, $3, $4)`,
		newShipment.DateOfDispatch,
		StringToJson(shippingAddress),
		newShipment.ShipmentStatus,
		newShipment.IsUrgent,
	).Scan(&shipmentId)

	newShipment.ShipmentID = shipmentId
	if err != nil {
		log.Fatal("Insert failed:", err)
		return nil, err
	}
	return newShipment, nil
}

func (s ShipmentService) UpdateShipment(shipment Models.Shipment, id int) error {
	commandTag, err := s.DB.Exec(
		context.Background(),
		`UPDATE "Shipments" SET "dateOfDispatch"=$1, "shippingAddress"=$2, "shipmentStatus"=$3, "isUrgent"=$4 WHERE "shipmentID"=$5`,
		shipment.DateOfDispatch,
		StringToJson(shipment.ShippingAddress),
		shipment.ShipmentStatus,
		shipment.IsUrgent,
		id,
	)

	if err != nil {
		log.Println("Update failed:", err)
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("no shipment found with given ID")
	}

	return nil
}

func (s ShipmentService) DeleteShipment(id int) error {
	commandTag, err := s.DB.Exec(
		context.Background(),
		`DELETE FROM "Shipments" WHERE "shipmentID"=$1`,
		id,
	)

	if err != nil {
		log.Println("Delete failed:", err)
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("no shipment found with given ID")
	}

	return nil
}

func (s ShipmentService) GetShipmentById(id int) (*Models.Shipment, error) {
	var shipment Models.Shipment
	err := s.DB.QueryRow(
		context.Background(),
		`SELECT * FROM "Shipments" WHERE "shipmentID"=$1`,
		id,
	).Scan(
		&shipment.ShipmentID,
		&shipment.DateOfDispatch,
		&shipment.ShippingAddress,
		&shipment.ShipmentStatus,
		&shipment.IsUrgent,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("shipment not found")
		}
		log.Println("Query failed:", err)
		return nil, err
	}

	return &shipment, nil
}

// util
func StringToJson(stringToChange string) []byte {
	shipmentAddressJSON, err := json.Marshal(stringToChange)
	if err != nil {
		log.Fatal(err)
	}
	return shipmentAddressJSON
}
