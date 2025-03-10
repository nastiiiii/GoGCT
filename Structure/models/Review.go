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
