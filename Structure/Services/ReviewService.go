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

type IReviewService interface {
	CreateComment(review Models.Review) (string, error)
	CreateReviewByParams(accountId int, performanceId int, reviewComment string, reviewRating int) Models.Review
	DeleteComment(commentId int)
	GetReviewsByPerformanceId(performanceId int) []Models.Review
	GetReviewsByAccountId(accountId int) []Models.Review
}

func (r *ReviewService) CreateComment(review Models.Review) (string, error) {
	query := `INSERT INTO reviews (account_id, performance_id, review_comment, review_rating, review_date)
	          VALUES ($1, $2, $3, $4, $5) RETURNING review_id`

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

func (r *ReviewService) CreateReviewByParams(accountId int, performanceId int, reviewComment string, reviewRating int) (Models.Review, error) {
	review := Models.Review{
		AccountId:     accountId,
		PerformanceId: performanceId,
		ReviewComment: reviewComment,
		ReviewRating:  reviewRating,
		ReviewDate:    time.Now(),
	}

	message, err := r.CreateComment(review)
	if err != nil {
		return Models.Review{}, err
	}

	log.Println(message)
	return review, nil
}

func (r *ReviewService) DeleteComment(reviewId int) error {
	query := `DELETE FROM reviews WHERE review_id = $1`
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

func (r *ReviewService) GetReviewsByPerformanceId(performanceId int) ([]Models.Review, error) {
	query := `SELECT review_id, account_id, performance_id, review_comment, review_rating, review_date
	          FROM reviews WHERE performance_id = $1`

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

func (r *ReviewService) GetReviewsByAccountId(accountId int) ([]Models.Review, error) {
	query := `SELECT review_id, account_id, performance_id, review_comment, review_rating, review_date
	          FROM reviews WHERE account_id = $1`

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
