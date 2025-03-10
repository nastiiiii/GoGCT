package models

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

// Warning!! you can change the signature of the function if you see needs to be changed, or you can add new one
const (
	Pending   BookingStatus = "Pending"
	Shipped   BookingStatus = "Shipped"
	Delivered BookingStatus = "Delivered"
)

type IShipmentService interface {
	CreateShipment(dateOfDispatch time.Time, shippingAddress string, shipmentStatus BookingStatus, isUrgent bool) (*Shipment, error)
	UpdateShipment(shipment Shipment, id int) Shipment
	DeleteShipment(id int)
	GetShipmentById(id int) Shipment
	GetPrice() float64
}

// @todo implement getters inside
type Shipment struct {
	shipmentID      int
	dateOfDispatch  time.Time
	shippingAddress string
	shipmentStatus  BookingStatus
	isUrgent        bool
	service         IShipmentService
}
type ShipmentService struct {
	DB *pgx.Conn
}

func (s ShipmentService) CreateShipment(dateOfDispatch time.Time, shippingAddress string, shipmentStatus BookingStatus, isUrgent bool) (*Shipment, error) {
	var err error
	newShipment := &Shipment{
		dateOfDispatch:  dateOfDispatch,
		shippingAddress: shippingAddress,
		shipmentStatus:  shipmentStatus,
		isUrgent:        isUrgent,
		service:         s,
	}
	var shipmentId int
	err = s.DB.QueryRow(
		context.Background(),
		`INSERT INTO "Shipments" ("dateOfDispatch", "shippingAddress", "shipmentStatus", "isUrgent") VALUES ($1, $2, $3, $4)`,
		newShipment.dateOfDispatch,
		StringToJson(shippingAddress),
		newShipment.shipmentStatus,
		newShipment.isUrgent,
	).Scan(&shipmentId)

	newShipment.shipmentID = shipmentId
	if err != nil {
		log.Fatal("Insert failed:", err)
		return nil, err
	}
	return newShipment, nil
}

func (s ShipmentService) UpdateShipment(shipment Shipment, id int) Shipment {
	//TODO implement me
	panic("implement me")
}

func (s ShipmentService) DeleteShipment(id int) {
	//TODO implement me
	panic("implement me")
}

func (s ShipmentService) GetShipmentById(id int) Shipment {
	//TODO implement me
	panic("implement me")
}

func (s ShipmentService) GetPrice() float64 {
	//TODO implement me
	panic("implement me")
}

// util
func StringToJson(stringToChange string) []byte {
	shipmentAddressJSON, err := json.Marshal(stringToChange)
	if err != nil {
		log.Fatal(err)
	}
	return shipmentAddressJSON
}
