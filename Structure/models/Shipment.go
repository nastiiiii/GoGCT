package models

import (
	"time"
)

// Warning!! you can change the signature of the function if you see needs to be changed, or you can add new one
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
