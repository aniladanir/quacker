package model

import "github.com/aniladanir/quacker/api/data"

type User struct {
	tableName      struct{} `pg:"users"`
	Id             int      `pg:",pk"`
	Tag            string   `pg:",unique"`
	Email          string   `pg:",unique"`
	Password       string   `pg:",notnull"`
	Name           string   `pg:",notnull"`
	Description    string
	Quacks         []*Quack `pg:"rel:has-many"`
	Likes          []*Like  `pg:"rel:has-many"`
	Followers      []User   `pg:"many2many:connections,join_fk:follower_id"`
	FollowerCount  int      `pg:"-"`
	FollowingCount int      `pg:"-"`
}

func (u User) IsModel() bool {
	return true
}

func (u *User) ToResponse() *data.UserResponse {
	response := &data.UserResponse{
		Tag:            u.Tag,
		Name:           u.Name,
		Description:    u.Description,
		FollowerCount:  u.FollowerCount,
		FollowingCount: u.FollowingCount,
	}

	return response
}
