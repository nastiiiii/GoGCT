package models

import (
	"time"
)

type Review struct {
	ReviewId      int
	AccountId     int
	PerformanceId int
	ReviewComment string
	ReviewRating  int
	ReviewDate    time.Time
}

func NewReview(accountId, performanceId int, review string, reviewRating int, reviewDate time.Time) *Review {
	return &Review{
		AccountId:     accountId,
		PerformanceId: performanceId,
		ReviewComment: review,
		ReviewRating:  reviewRating,
		ReviewDate:    reviewDate,
	}
}
