package Controller

import (
	"GCT/Structure/Services"
	"GCT/Structure/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ReviewController struct {
	ReviewService Services.IReviewService
}

func NewReviewController(reviewService Services.IReviewService) *ReviewController {
	return &ReviewController{ReviewService: reviewService}
}

func (rc *ReviewController) CreateReview(c *gin.Context) {
	var request struct {
		AccountId     int    `json:"accountId"`
		PerformanceId int    `json:"performanceId"`
		ReviewComment string `json:"reviewComment"`
		ReviewRating  int    `json:"reviewRating"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review := models.Review{
		AccountId:     request.AccountId,
		PerformanceId: request.PerformanceId,
		ReviewComment: request.ReviewComment,
		ReviewRating:  request.ReviewRating,
		ReviewDate:    time.Now(),
	}

	message, err := rc.ReviewService.CreateReview(review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (rc *ReviewController) DeleteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("reviewId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = rc.ReviewService.DeleteReview(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
}

func (rc *ReviewController) GetReviewsByPerformanceId(c *gin.Context) {
	performanceId, err := strconv.Atoi(c.Param("performanceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reviews, err := rc.ReviewService.GetReviewsByPerformanceId(performanceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func (rc *ReviewController) GetReviewsByAccountId(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reviews, err := rc.ReviewService.GetReviewsByAccountId(accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func SetupReviewRouters(router *gin.Engine, service Services.IReviewService) {
	controller := NewReviewController(service)
	reviewRoutes := router.Group("/review")
	{
		reviewRoutes.POST("/create", controller.CreateReview)
		reviewRoutes.DELETE("/:reviewId", controller.DeleteReview)
		reviewRoutes.GET("/performance/:performanceId", controller.GetReviewsByPerformanceId)
		reviewRoutes.GET("/account/:accountId", controller.GetReviewsByAccountId)
	}
}
