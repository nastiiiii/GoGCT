package Controller

import (
	"GCT/Structure/Services"
	"GCT/Structure/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransactionController struct {
	service Services.ITransactionService
}

func NewTransactionController(service Services.ITransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := tc.service.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (tc *TransactionController) GetTransactionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tram, err := tc.service.GetTransactionById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tram)
}

func (tc *TransactionController) GetTransactionByAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transactions, err := tc.service.GetTransactionByAccount(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
func (tc *TransactionController) UpdateTransaction(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTransaction, err := tc.service.UpdateTransaction(transactionId, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTransaction)
}

func (tc *TransactionController) DeleteTransaction(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	if !tc.service.DeleteTransaction(transactionId) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete transaction"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

func (tc *TransactionController) GetTransactionsByStatus(c *gin.Context) {
	status := c.Param("status")
	transactions := tc.service.GetTransactionsByStatus(models.TransactionStatus(status))
	c.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) GetByConfirmationId(c *gin.Context) {
	confirmationId := c.Param("confirmationId")
	transaction, err := tc.service.GetByConfirmationId(confirmationId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) ProcessTransactionPayment(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	if err := tc.service.ProcessTransactionPayment(transactionId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction payment processed successfully"})
}

func SetupTransactionRoutes(router *gin.Engine, service Services.ITransactionService) {
	tc := NewTransactionController(service)
	routes := router.Group("/transactions")
	{
		routes.POST("/", tc.CreateTransaction)
		routes.GET("/:id", tc.GetTransactionById)
		routes.GET("/account/:accountId", tc.GetTransactionByAccount)
		routes.GET("/status/:status", tc.GetTransactionsByStatus)
		routes.GET("/confirmation/:confirmationId", tc.GetByConfirmationId)
		routes.PUT("/:id", tc.UpdateTransaction)
		routes.DELETE("/:id", tc.DeleteTransaction)
		routes.POST("/:id/payment", tc.ProcessTransactionPayment)
	}
}
