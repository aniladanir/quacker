package data

import "github.com/aniladanir/quacker/quackerDb/model"

// Use for post, put operations
type UserEntry struct {
	Tag         string
	Email       string
	Password    string
	Name        string
	Description string
}

func (u *UserEntry) ToModel() *model.User {
	user := &model.User{
		Tag:         u.Tag,
		Email:       u.Tag,
		Password:    u.Tag,
		Name:        u.Tag,
		Description: u.Tag,
	}

	return user
}
