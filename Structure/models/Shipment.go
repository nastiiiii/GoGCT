package models

import (
	"time"
)

// represents shipment status
const (
	Pending   BookingStatus = "Pending"
	Shipped   BookingStatus = "Shipped"
	Delivered BookingStatus = "Delivered"
)

type Shipment struct {
	ShipmentID      int
	DateOfDispatch  time.Time
	ShippingAddress string
	ShipmentStatus  BookingStatus
	IsUrgent        bool
}

func NewShipment(dateOfDispatch time.Time, shippingAddress string, shipmentStatus BookingStatus, isUrgent bool) *Shipment {
	return &Shipment{
		DateOfDispatch:  dateOfDispatch,
		ShippingAddress: shippingAddress,
		ShipmentStatus:  shipmentStatus,
		IsUrgent:        isUrgent,
	}
}
