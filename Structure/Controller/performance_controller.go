package Controller

import (
	"GCT/Structure/Interfaces"
	Models "GCT/Structure/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PerformanceController struct {
	service Interfaces.IPerformanceService
}

func NewPerformanceController(service Interfaces.IPerformanceService) *PerformanceController {
	return &PerformanceController{service}
}

func (pc *PerformanceController) CreatePerformance(c *gin.Context) {
	var performance Models.Performance
	if err := c.ShouldBindJSON(&performance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := pc.service.CreatePerformance(performance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (pc *PerformanceController) UpdatePerformance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var performance Models.Performance
	if err := c.ShouldBindJSON(&performance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatePerformance, err := pc.service.UpdatePerformance(performance, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatePerformance)
}

func (pc *PerformanceController) DeletePerformance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if success := pc.service.DeletePerformance(id); !success {
		c.JSON(http.StatusOK, gin.H{"success": true})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Performance has been deleted"})
}

func (pc *PerformanceController) GetPerformanceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	performance, err := pc.service.GetPerformanceById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, performance)
}

func (pc *PerformanceController) GetPerformanceByName(c *gin.Context) {
	performance, err := pc.service.GetPerformanceByName(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, performance)
}

func (pc *PerformanceController) GetPerformances(c *gin.Context) {
	performances := pc.service.GetPerformances()
	c.JSON(http.StatusOK, performances)
}

func (pc *PerformanceController) GetAllSeats(c *gin.Context) {
	var request struct {
		PerformanceName string `json:"performanceName"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	performance, err := pc.service.GetPerformanceByName(request.PerformanceName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Performance not found"})
		return
	}

	seats, err := pc.service.GetAllSeats(performance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seats)
}

func (pc *PerformanceController) GetAvailableSeats(c *gin.Context) {
	var request struct {
		PerformanceName string `json:"performanceName"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	performance, err := pc.service.GetPerformanceByName(request.PerformanceName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Performance not found"})
		return
	}

	seats, err := pc.service.GetAvailableSeats(performance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seats)
}

func (pc *PerformanceController) GetSeatPrice(c *gin.Context) {
	var request struct {
		PerformanceName string `json:"performanceName"`
		SeatBand        string `json:"seatBand"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	performance, err := pc.service.GetPerformanceByName(request.PerformanceName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Performance not found"})
		return
	}

	price, err := pc.service.GetSeatPrice(performance, request.SeatBand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"seatPrice": price})
}

func (pc *PerformanceController) ChangeSeatAvailability(c *gin.Context) {
	var request struct {
		PerformanceName string `json:"performanceName"`
		SeatBand        string `json:"seatBand"`
		SeatNumber      int    `json:"seatNumber"`
		Availability    bool   `json:"availability"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	performance, err := pc.service.GetPerformanceByName(request.PerformanceName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Performance not found"})
		return
	}

	success, err := pc.service.ChangeSeatAvailability(&performance, request.SeatBand, request.SeatNumber, request.Availability)
	if err != nil || !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Seat availability updated successfully"})

}

func SetUpPerformanceRouters(router *gin.Engine, service Interfaces.IPerformanceService) {
	controller := NewPerformanceController(service)
	performanceRoutes := router.Group("/performance")
	{
		performanceRoutes.POST("/create", controller.CreatePerformance)
		performanceRoutes.PUT("/:id", controller.UpdatePerformance)
		performanceRoutes.DELETE("/:id", controller.DeletePerformance)
		performanceRoutes.GET("/:id", controller.GetPerformanceById)
		performanceRoutes.GET("/name/:name", controller.GetPerformanceByName)
		performanceRoutes.GET("/", controller.GetPerformances)
		performanceRoutes.POST("/seats/all", controller.GetAllSeats)
		performanceRoutes.POST("/seats/available", controller.GetAvailableSeats)
		performanceRoutes.POST("/seats/price", controller.GetSeatPrice)
		performanceRoutes.POST("/seats/update", controller.ChangeSeatAvailability)

	}
}
