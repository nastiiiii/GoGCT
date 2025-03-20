package Controller

import (
	"GCT/Structure/Interfaces"
	Models "GCT/Structure/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TicketController Description: Manages ticket creation, update, get and delete
type TicketController struct {
	service Interfaces.ITicketService
}

func NewTicketController(service Interfaces.ITicketService) *TicketController {
	return &TicketController{service: service}
}

func (tc *TicketController) CreateTicket(c *gin.Context) {
	var ticket Models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTicket, err := tc.service.CreateTicket(ticket)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTicket)
}

func (tc *TicketController) GetTicketById(c *gin.Context) {
	ticketId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket, err := tc.service.GetTicketById(ticketId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (tc *TicketController) GetTicketsByTransactionId(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tickets, err := tc.service.GetTicketsByTransactionId(transactionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (tc *TicketController) GetPriceByTicket(c *gin.Context) {
	var ticket Models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	price := tc.service.GetPriceByTicket(ticket)
	c.JSON(http.StatusOK, gin.H{"price": price})
}

func (tc *TicketController) GetTicketsPriceByTransaction(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("transactionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	price := tc.service.GetTicketsPriceByTransaction(transactionId)
	c.JSON(http.StatusOK, gin.H{"total_price": price})
}

func (tc *TicketController) GetTicketsByAccount(c *gin.Context) {
	token := c.Query("token")
	tickets, err := tc.service.GetTicketsByAccount(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (tc *TicketController) UpdateTicket(c *gin.Context) {
	ticketId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ticket Models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTicket, err := tc.service.UpdateTicket(ticketId, ticket)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTicket)
}

func (tc *TicketController) DeleteTicket(c *gin.Context) {
	ticketId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !tc.service.DeleteTicket(ticketId) {
		c.JSON(http.StatusOK, gin.H{"message": "Ticket not found"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Ticket deleted"})
}

func (tc *TicketController) DeleteTicketsByTransactionId(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !tc.service.DeleteTicketsByTransactionId(transactionId) {
		c.JSON(http.StatusOK, gin.H{"message": "Ticket not found"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Ticket deleted"})
}

// SetupTicketRoutes setting up the router
func SetupTicketRoutes(router *gin.Engine, service Interfaces.ITicketService) {
	tc := NewTicketController(service)
	tickets := router.Group("/tickets")
	{
		tickets.POST("/", tc.CreateTicket)
		tickets.GET("/:id", tc.GetTicketById)
		tickets.GET("/transaction/:transactionId", tc.GetTicketsByTransactionId)
		tickets.GET("/price", tc.GetPriceByTicket)
		tickets.GET("/transaction/:transactionId/price", tc.GetTicketsPriceByTransaction)
		tickets.GET("/account", tc.GetTicketsByAccount)
		tickets.PUT("/:id", tc.UpdateTicket)
		tickets.DELETE("/:id", tc.DeleteTicket)
		tickets.DELETE("/transaction/:transactionId", tc.DeleteTicketsByTransactionId)
	}
}
