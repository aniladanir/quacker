package model

type Quack struct {
	Id        int `pg:",pk"`
	Content   string
	Quacks    []*Quack    `pg:"rel:has-many"`
	Favorites []*Favorite `pg:"rel:has-many"`
	UserId    int         `pg:",notnull"`
}

func (q Quack) IsModel() bool {
	return true
}
