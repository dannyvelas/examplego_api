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

	rows, err := reviewsRepo.database.driver.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("review_repo: GetActive: %v", NewError(ErrDatabaseQuery, err))
	}
	defer rows.Close()

	reviews := []models.Review{}
	for rows.Next() {
		var review models.Review

		err := rows.Scan(
			&review.Id,
			&review.UserId,
			&review.BookId,
			&review.TitleAndAuthor,
			&review.ReviewDate,
			&review.AmtStars,
			&review.Description,
			&review.IsAnonymous,
		)
		if err != nil {
			return nil, fmt.Errorf("review_repo: GetActive: %v", NewError(ErrScanningRow, err))
		}

		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("review_repo: GetActive: %v", NewError(ErrIterating, err))
	}

	return reviews, nil
}

func (reviewsRepo *ReviewsRepo) GetAll(limit, offset uint) ([]models.Review, error) {
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

	rows, err := reviewsRepo.database.driver.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("review_repo: GetAll: %v", NewError(ErrDatabaseQuery, err))
	}
	defer rows.Close()

	reviews := []models.Review{}
	for rows.Next() {
		var review models.Review
		err := rows.Scan(
			&review.Id,
			&review.UserId,
			&review.BookId,
			&review.TitleAndAuthor,
			&review.ReviewDate,
			&review.AmtStars,
			&review.Description,
			&review.IsAnonymous,
		)

		if err != nil {
			return nil, fmt.Errorf("review_repo: GetAll: %v", NewError(ErrScanningRow, err))
		}

		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("review_repo: GetAll: %v", NewError(ErrIterating, err))
	}

	return reviews, nil
}
