package models

import (
	"github.com/jackc/pgx/v5"
	"time"
)

type Review struct {
	reviewId      int
	accountId     int
	performanceId int
	reviewComment string
	reviewRating  int
	reviewDate    time.Time
	service       IReviewService
}

type ReviewService struct {
	DB *pgx.Conn
}

type IReviewService interface {
	IsConfirmed(account Account, performance Performance) bool
	CreateComment(review Review)
	CreateReviewByParams(accountId int, performanceId int, reviewComment string, reviewRating int) Review
	DeleteComment(commentId int)
	GetReviewsByPerformanceId(performanceId int) []Review
	GetReviewsByAccountId(accountId int) []Review
}

func (r ReviewService) IsConfirmed(account Account, performance Performance) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) CreateComment(review Review) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) CreateReviewByParams(accountId int, performanceId int, reviewComment string, reviewRating int) Review {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) DeleteComment(commentId int) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) GetReviewsByPerformanceId(performanceId int) []Review {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) GetReviewsByAccountId(accountId int) []Review {
	//TODO implement me
	panic("implement me")
}
