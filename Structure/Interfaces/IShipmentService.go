package Interfaces

import (
	Models "GCT/Structure/models"
	"time"
)

type IShipmentService interface {
	CreateShipmentByParams(dateOfDispatch time.Time, shippingAddress string, shipmentStatus Models.BookingStatus, isUrgent bool) (*Models.Shipment, error)
	CreateShipment(shipment Models.Shipment) (*Models.Shipment, error)
	UpdateShipment(shipment Models.Shipment, id int) (bool, error)
	DeleteShipment(id int) error
	GetShipmentById(id int) (*Models.Shipment, error)
}
