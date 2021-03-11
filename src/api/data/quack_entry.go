package data

import "github.com/aniladanir/quacker/quackerDb/model"

//Use for post operation
type QuackEntry struct {
	QuackTypeId int
	Content     string
	UserId      int
	ParentId    int
}

func (q *QuackEntry) ToModel() *model.Quack {
	quack := &model.Quack{
		QuackTypeId: q.QuackTypeId,
		Content:     q.Content,
		UserId:      q.UserId,
		ParentId:    q.ParentId,
	}

	return quack
}
