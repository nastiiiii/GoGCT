package Controller

import (
	"GCT/Structure/Interfaces"
	"GCT/Structure/Services"
	"GCT/Structure/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DiscountController is responsible for managing discount operations
type DiscountController struct {
	DiscountService Interfaces.IDiscountService
}

func NewDiscountController(service Interfaces.IDiscountService) *DiscountController {
	return &DiscountController{DiscountService: service}
}

func (controller *DiscountController) GetDiscounts(c *gin.Context) {
	discounts, err := controller.DiscountService.LoadDiscounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load discounts"})
		return
	}
	c.JSON(http.StatusOK, discounts)
}

func (controller *DiscountController) SaveDiscount(c *gin.Context) {
	var discount models.Discount
	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := controller.DiscountService.SaveDiscount(discount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save discount"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Discount saved successfully"})
}

func (controller *DiscountController) DeleteDiscount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := controller.DiscountService.DeleteDiscount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete discount"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discount deleted successfully"})
}

func (controller *DiscountController) ApplyBestDiscount(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction data"})
		return
	}

	controller.DiscountService.ApplyBestDiscount(&transaction)
	c.JSON(http.StatusOK, transaction) // Respond with the updated transaction
}

// SetUpDiscountRouters setting up the router
func SetUpDiscountRoutes(router *gin.Engine, service Services.DiscountService) {
	controller := NewDiscountController(&service)
	discountRoutes := router.Group("/discount")
	{
		discountRoutes.GET("/discounts", controller.GetDiscounts)
		discountRoutes.POST("/discounts", controller.SaveDiscount)
		discountRoutes.DELETE("/discounts/:id", controller.DeleteDiscount)
		discountRoutes.POST("/discounts/apply", controller.ApplyBestDiscount)
	}
}
