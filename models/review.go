package models

type Review struct {
	Id            int
	UserId    string
	BookId  string
	TitleAndAuthor string
	ReviewDate     int
	AmtStars       int
	Description   int
	IsAnonymous   bool
}
