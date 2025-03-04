package models

import "time"

type Review struct {
	//reviewId      int
	accountId     int
	performanceId int
	reviewComment string
	reviewRating  int
	reviewDate    time.Time
}

type IReviewService interface {
	isConfirmed(account Account, performance Performance) bool
	createComment(review Review)
}

func (r Review) isConfirmed(account Account, performance Performance) {

}

func (r Review) createComment(review Review) {

}
