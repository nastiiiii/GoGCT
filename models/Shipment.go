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
