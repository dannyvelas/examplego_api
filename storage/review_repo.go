package storage

import (
	"github.com/dannyvelas/go-backend/models"
)

type ReviewRepo struct {
	database Database
}

func NewReviewRepo(database Database) ReviewRepo {
	return ReviewRepo{database: database}
}

func (reviewRepo ReviewRepo) GetActive(limit, offset int) ([]models.Review, error) {
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

	rows, err := reviewRepo.database.driver.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
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
			return nil, err
		}

		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (reviewRepo *ReviewRepo) GetAll(limit, offset int) ([]models.Review, error) {
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

	rows, err := reviewRepo.database.driver.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
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
			return nil, err
		}

		reviews = append(reviews, review)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}
