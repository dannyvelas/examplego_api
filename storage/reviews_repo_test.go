package storage

import (
	"github.com/dannyvelas/examplego_api/config"
	"github.com/dannyvelas/examplego_api/models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"testing"
)

type reviewsRepoSuite struct {
	suite.Suite
	reviewsRepo ReviewsRepo
	migrator    migrate.Migrate
}

func (suite *reviewsRepoSuite) SetupSuite() {
	config := config.NewConfig()

	database, err := NewDatabase(config.Postgres())
	if err != nil {
		log.Fatal().Msgf("Failed to start database: %v", err)
		return
	}

	migrator, err := GetMigrator(database)
	if err != nil {
		log.Fatal().Msgf("Failed to get migrator: %v", err)
	}

	suite.reviewsRepo = NewReviewsRepo(database)
	suite.migrator = migrator
}

func (suite reviewsRepoSuite) TearDownTest() {
	suite.reviewsRepo.deleteAll()
}

func (suite reviewsRepoSuite) TearDownSuite() {
	suite.migrator.Down()
}

func (suite reviewsRepoSuite) TestGetAllReviews_EmptySlice_Positive() {
	reviews, err := suite.reviewsRepo.GetAll(defaultLimit, defaultOffset)
	suite.NoError(err, "no error when getting all reviews when the table is empty")
	suite.Equal(len(reviews), 0, "length of reviews should be 0, since it is empty slice")
	suite.Equal(reviews, []models.Review(nil), "reviews is an empty slice")
}

func TestReviewsRepo(t *testing.T) {
	suite.Run(t, new(reviewsRepoSuite))
}
