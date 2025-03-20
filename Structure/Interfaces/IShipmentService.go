package Interfaces

import (
	Models "GCT/Structure/models"
	"time"
)

// IShipmentService Description: Shows all methods which ShipmentService is implemented and to which the controllers have access
type IShipmentService interface {
	//Create
	CreateShipmentByParams(dateOfDispatch time.Time, shippingAddress string, shipmentStatus Models.BookingStatus, isUrgent bool) (*Models.Shipment, error)
	CreateShipment(shipment Models.Shipment) (*Models.Shipment, error)
	//Update
	UpdateShipment(shipment Models.Shipment, id int) (bool, error)
	//Delete
	DeleteShipment(id int) error
	//Get
	GetShipmentById(id int) (*Models.Shipment, error)
}
