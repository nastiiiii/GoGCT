package models

import "time"

const (
	Pending   BookingStatus = "Pending"
	Shipped   BookingStatus = "Shipped"
	Delivered BookingStatus = "Delivered"
)

type Shipment struct {
	//shipmentID      int
	DateOfDispatch  time.Time
	ShippingAddress string
	ShipmentStatus  BookingStatus
	IsUrgent        bool
}

type IShipmentService interface {
	createShipment(shipment Shipment) Shipment
	updateShipment(shipment Shipment, id int) Shipment
	deleteShipment(id int)
	getShipmentById(id int) Shipment
	getPrice() float64
}
