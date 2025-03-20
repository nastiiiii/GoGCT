package Services

import (
	Models "GCT/Structure/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

type ReviewService struct {
	DB *pgx.Conn
}

// Approve
func (r ReviewService) CreateReview(review Models.Review) (string, error) {
	query := `INSERT INTO "Reviews" ("accountID", "performanceID", "reviewComment", "reviewRating", "reviewDate")
	          VALUES ($1, $2, $3, $4, $5) RETURNING "reviewID"`

	err := r.DB.QueryRow(
		context.Background(),
		query,
		review.AccountId,
		review.PerformanceId,
		review.ReviewComment,
		review.ReviewRating,
		time.Now(),
	).Scan(&review.ReviewId)

	if err != nil {
		return "", errors.New("failed to create review")
	}

	return "Review successfully created", nil
}

// Approve
func (r ReviewService) CreateReviewByParams(accountId int, performanceId int, reviewComment string, reviewRating int) (Models.Review, error) {
	review := Models.Review{
		AccountId:     accountId,
		PerformanceId: performanceId,
		ReviewComment: reviewComment,
		ReviewRating:  reviewRating,
		ReviewDate:    time.Now(),
	}

	message, err := r.CreateReview(review)
	if err != nil {
		return Models.Review{}, err
	}

	log.Println(message)
	return review, nil
}

// Approve
func (r ReviewService) DeleteReview(reviewId int) error {
	query := `DELETE FROM "Reviews" WHERE "reviewID" = $1`
	result, err := r.DB.Exec(context.Background(), query, reviewId)

	if err != nil {
		return errors.New("failed to delete review")
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no review found with the given ID")
	}

	return nil
}

// Aprrove
func (r ReviewService) GetReviewsByPerformanceId(performanceId int) ([]Models.Review, error) {
	query := `SELECT "reviewID", "accountID", "performanceID", "reviewComment", "reviewRating", "reviewDate"
	          FROM "Reviews" WHERE "performanceID" = $1`

	rows, err := r.DB.Query(context.Background(), query, performanceId)
	if err != nil {
		return nil, errors.New("could not retrieve reviews")
	}
	defer rows.Close()

	var reviews []Models.Review
	for rows.Next() {
		var review Models.Review
		err := rows.Scan(&review.ReviewId, &review.AccountId, &review.PerformanceId,
			&review.ReviewComment, &review.ReviewRating, &review.ReviewDate)
		if err != nil {
			log.Println("Error scanning review:", err)
			continue
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// Approve
func (r ReviewService) GetReviewsByAccountId(accountId int) ([]Models.Review, error) {
	query := `SELECT "reviewID", "accountID", "performanceID", "reviewComment", "reviewRating", "reviewDate"
	          FROM "Reviews" WHERE "accountID" = $1`

	rows, err := r.DB.Query(context.Background(), query, accountId)
	if err != nil {
		return nil, errors.New("could not retrieve reviews")
	}
	defer rows.Close()

	var reviews []Models.Review
	for rows.Next() {
		var review Models.Review
		err := rows.Scan(&review.ReviewId, &review.AccountId, &review.PerformanceId,
			&review.ReviewComment, &review.ReviewRating, &review.ReviewDate)
		if err != nil {
			log.Println("Error scanning review:", err)
			continue
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
