package models

import (
	"time"
)

type Review struct {
	Id          int
	UserId      string
	BookId      string
	ReviewDate  time.Time
	AmtStars    int
	Description int
}
