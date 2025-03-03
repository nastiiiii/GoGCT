package models

import "time"

const (
	Planned BookingStatus = iota
	Dispatched
	Delivered
	Canceled
)

type Shipment struct {
	//shipmentID      int
	dateOfDispatch  time.Time
	shippingAddress string
	shipmentStatus  BookingStatus
	isUrgent        bool
}

type IShipmentService interface {
	createShipment(shipment Shipment) Shipment
	updateShipment(shipment Shipment, id int) Shipment
	deleteShipment(id int)
	getShipmentById(id int) Shipment
	getPrice() float64
}
