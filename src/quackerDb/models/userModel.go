package models

type User struct {
	Id          int
	Tag         string `pg:",unique"`
	Email       string `pg:",unique"`
	Password    string `pg:",notnull"`
	Name        string `pg:",notnull"`
	Description string
	Quacks      []*Quack    `pg:"rel:has-many"`
	Favorites   []*Favorite `pg:"rel:has-many"`
	Followers   []User      `pg:"many2many:user_to_users,join_fk:follower_id"`
}

func (u User) IsModel() bool {
	return true
}
