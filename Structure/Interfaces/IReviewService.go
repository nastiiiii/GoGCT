package Interfaces

import Models "GCT/Structure/models"

type IReviewService interface {
	CreateReview(review Models.Review) (string, error)
	CreateReviewByParams(accountId int, performanceId int, reviewComment string, reviewRating int) (Models.Review, error)
	DeleteReview(reviewId int) error
	GetReviewsByPerformanceId(performanceId int) ([]Models.Review, error)
	GetReviewsByAccountId(accountId int) ([]Models.Review, error)
}
