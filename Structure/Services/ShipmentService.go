package Services

import (
	Models "GCT/Structure/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

// ShipmentService implements the database operations and businesses logic related to Shipment
type ShipmentService struct {
	DB *pgx.Conn
}

func (s ShipmentService) CreateShipmentByParams(dateOfDispatch time.Time, shippingAddress string, shipmentStatus Models.BookingStatus, isUrgent bool) (*Models.Shipment, error) {
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
		newShipment.ShippingAddress,
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

func (s ShipmentService) CreateShipment(shipment Models.Shipment) (*Models.Shipment, error) {
	var err error
	var shipmentId int
	query := `INSERT INTO "Shipments" ("dateOfDispatch", "shippingAddress", "shipmentStatus", "isUrgent")
	          VALUES ($1, $2, $3, $4) RETURNING "shipmentID"`

	err = s.DB.QueryRow(
		context.Background(),
		query,
		shipment.DateOfDispatch,
		shipment.ShippingAddress,
		shipment.ShipmentStatus,
		shipment.IsUrgent,
	).Scan(&shipmentId)
	shipment.ShipmentID = shipmentId
	if err != nil {
		log.Fatal("Insert failed:", err)
		return nil, err
	}
	return &shipment, nil
}

func (s ShipmentService) UpdateShipment(shipment Models.Shipment, id int) (bool, error) {
	commandTag, err := s.DB.Exec(
		context.Background(),
		`UPDATE "Shipments" SET "dateOfDispatch"=$1, "shippingAddress"=$2, "shipmentStatus"=$3, "isUrgent"=$4 WHERE "shipmentID"=$5`,
		shipment.DateOfDispatch,
		shipment.ShippingAddress,
		shipment.ShipmentStatus,
		shipment.IsUrgent,
		id,
	)

	if err != nil {
		log.Println("Update failed:", err)
		return false, err
	}

	if commandTag.RowsAffected() == 0 {
		return false, errors.New("no shipment found with given ID")
	}

	return true, nil
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
		`SELECT "dateOfDispatch", "shippingAddress", "shipmentStatus", "isUrgent" FROM "Shipments" WHERE "shipmentID"=$1`,
		id,
	).Scan(
		&shipment.DateOfDispatch,
		&shipment.ShippingAddress,
		&shipment.ShipmentStatus,
		&shipment.IsUrgent,
	)
	shipment.ShipmentID = id

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("shipment not found")
		}
		log.Println("Query failed:", err)
		return nil, err
	}

	return &shipment, nil
}
