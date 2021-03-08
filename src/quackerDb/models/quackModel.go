package models

type Quack struct {
	Id        int
	Content   string
	Quacks    []*Quack    `pg:"rel:has-many"`
	Favorites []*Favorite `pg:"rel:has-many"`
	UserId    int
}

func (q Quack) IsModel() bool {
	return true
}
