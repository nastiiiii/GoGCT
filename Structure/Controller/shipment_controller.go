package Controller

import (
	"GCT/Structure/Services"
	"GCT/Structure/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ShipmentController struct {
	shipmentService Services.IShipmentService
}

func NewShipmentController(service Services.IShipmentService) *ShipmentController {
	return &ShipmentController{
		service,
	}
}

func (sc *ShipmentController) CreateShipment(c *gin.Context) {
	var request struct {
		DateOfDispatch  string `json:"dateOfDispatch"`
		ShippingAddress string `json:"shippingAddress"`
		ShipmentStatus  string `json:"shipmentStatus"`
		IsUrgent        bool   `json:"isUrgent"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Convert date string to time.Time
	parsedDate, err := time.Parse("2006-01-02", request.DateOfDispatch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	shipment := models.Shipment{
		DateOfDispatch:  parsedDate,
		ShippingAddress: request.ShippingAddress,
		ShipmentStatus:  models.BookingStatus(request.ShipmentStatus),
		IsUrgent:        request.IsUrgent,
	}

	createdShipment, err := sc.shipmentService.CreateShipment(shipment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdShipment)
}

func (sc *ShipmentController) UpdateShipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shipment ID"})
		return
	}

	var request struct {
		DateOfDispatch  string `json:"dateOfDispatch"`
		ShippingAddress string `json:"shippingAddress"`
		ShipmentStatus  string `json:"shipmentStatus"`
		IsUrgent        bool   `json:"isUrgent"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", request.DateOfDispatch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	shipment := models.Shipment{
		ShipmentID:      id,
		DateOfDispatch:  parsedDate,
		ShippingAddress: request.ShippingAddress,
		ShipmentStatus:  models.BookingStatus(request.ShipmentStatus),
		IsUrgent:        request.IsUrgent,
	}

	success, err := sc.shipmentService.UpdateShipment(shipment, id)
	if err != nil || !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shipment updated successfully"})
}

func (sc *ShipmentController) DeleteShipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shipment ID"})
		return
	}

	if err := sc.shipmentService.DeleteShipment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shipment deleted successfully"})
}

func (sc *ShipmentController) GetShipmentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shipment ID"})
		return
	}

	shipment, err := sc.shipmentService.GetShipmentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipment)
}

func SetupShipmentRoutes(router *gin.Engine, service Services.IShipmentService) {
	controller := NewShipmentController(service)

	shipmentRoutes := router.Group("/shipments")
	{
		shipmentRoutes.POST("/create", controller.CreateShipment)
		shipmentRoutes.PUT("/:id", controller.UpdateShipment)
		shipmentRoutes.DELETE("/:id", controller.DeleteShipment)
		shipmentRoutes.GET("/:id", controller.GetShipmentById)
	}
}
