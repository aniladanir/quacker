package model

import (
	"time"

	"github.com/aniladanir/quacker/api/data"
)

type Quack struct {
	tableName   struct{} `pg:"quacks"`
	Id          int      `pg:",pk"`
	QuackTypeId int      `pg:",notnull"`
	Content     string
	DateCreated time.Time
	Quacks      []Quack `pg:"rel:has-many,join_fk:parent_id"`
	Likes       []*Like `pg:"rel:has-many"`
	ParentId    int
	UserId      int     `pg:",notnull"`
	UserTag		string	`pg:"-"`
	UserName	string	`pg:"-"`
	Parents		[]Quack	`pg:"-"`
}

func (q Quack) IsModel() bool {
	return true
}

func (q *Quack) ToResponse() *data.QuackResponse{
	response := &data.QuackResponse{
		Id: q.Id,
		QuackType: q.QuackTypeId,
		Content: q.Content,
		DateCreated: q.DateCreated,
		OwnerTag: q.UserTag,
		OwnerName: ,
	}
}
