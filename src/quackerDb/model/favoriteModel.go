package model

type Favorite struct {
	UserId  int `pg:"on_delete:RESTRICT"`
	QuackId int `pg:"on_delete:RESTRICT"`
}

func (f Favorite) IsModel() bool {

	return true
}
