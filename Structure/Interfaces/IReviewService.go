package Interfaces

import Models "GCT/Structure/models"

// IReviewService Description: Shows all methods which ReviewService is implemented and to which the controllers have access
type IReviewService interface {
	//Create
	CreateReview(review Models.Review) (string, error)
	CreateReviewByParams(accountId int, performanceId int, reviewComment string, reviewRating int) (Models.Review, error)
	//Delete
	DeleteReview(reviewId int) error
	//Get
	GetReviewsByPerformanceId(performanceId int) ([]Models.Review, error)
	GetReviewsByAccountId(accountId int) ([]Models.Review, error)
}
