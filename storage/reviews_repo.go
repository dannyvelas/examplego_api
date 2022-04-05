package storage

import (
	"fmt"
	"github.com/dannyvelas/examplego_api/models"
)

type ReviewsRepo struct {
	database Database
}

func NewReviewsRepo(database Database) ReviewsRepo {
	return ReviewsRepo{database: database}
}

func (reviewsRepo ReviewsRepo) GetActive(limit, offset uint) ([]models.Review, error) {
	const query = `
    SELECT
      reviews.id,
      reviews.user_id,
      reviews.book_id,
      books.title_and_author,
      reviews.review_date,
      reviews.amt_stars,
      reviews.description,
      reviews.is_anonymous
    FROM reviews
    LEFT JOIN books ON
      reviews.book_id = books.book_id 
    WHERE
      reviews.review_date <= EXTRACT(epoch FROM NOW())
      AND reviews.is_anonymous = FALSE
    LIMIT $1
    OFFSET $2
  `

	boundedLimit := getBoundedLimit(limit)
	rows, err := reviewsRepo.database.driver.Query(query, boundedLimit, offset)
	if err != nil {
		return nil, fmt.Errorf("reviews_repo: GetActive: %v", newError(ErrDatabaseQuery, err))
	}
	defer rows.Close()

	reviews := []models.Review{}
	for rows.Next() {
		var review models.Review

		err := rows.Scan(
			&review.Id,
			&review.UserId,
			&review.BookId,
			&review.ReviewDate,
			&review.AmtStars,
			&review.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("reviews_repo: GetActive: %v", newError(ErrScanningRow, err))
		}

		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("reviews_repo: GetActive: %v", newError(ErrIterating, err))
	}

	return reviews, nil
}

func (reviewsRepo ReviewsRepo) GetAll(limit, offset uint) ([]models.Review, error) {
	const query = `
    SELECT
      reviews.id,
      reviews.user_id,
      reviews.book_id,
      books.title_and_author,
      reviews.review_date,
      reviews.amt_stars,
      reviews.description,
      reviews.is_anonymous
    FROM reviews
    LEFT JOIN books ON
      reviews.book_id = books.book_id 
    LIMIT $1
    OFFSET $2
  `

	boundedLimit := getBoundedLimit(limit)
	rows, err := reviewsRepo.database.driver.Query(query, boundedLimit, offset)
	if err != nil {
		return nil, fmt.Errorf("reviews_repo: GetAll: %v", newError(ErrDatabaseQuery, err))
	}
	defer rows.Close()

	reviews := []models.Review{}
	for rows.Next() {
		var review models.Review
		err := rows.Scan(
			&review.Id,
			&review.UserId,
			&review.BookId,
			&review.ReviewDate,
			&review.AmtStars,
			&review.Description,
		)

		if err != nil {
			return nil, fmt.Errorf("reviews_repo: GetAll: %v", newError(ErrScanningRow, err))
		}

		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("reviews_repo: GetAll: %v", newError(ErrIterating, err))
	}

	return reviews, nil
}

func (reviewsRepo ReviewsRepo) deleteAll() (int64, error) {
	query := "DELETE FROM reviews"
	res, err := reviewsRepo.database.driver.Exec(query)
	if err != nil {
		return 0, fmt.Errorf("reviews_repo: deleteAll: %v", newError(ErrDatabaseExec, err))
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("reviews_repo: deleteAll: %v", newError(ErrGetRowsAffected, err))
	}

	return rowsAffected, nil
}
