package models

import (
	"time"
)

type Favorite struct {
	UserId      int
	QuackId     int
	DateCreated time.Time `pg:"default:now()"`
}

func (f Favorite) IsModel() bool {

	return true
}
