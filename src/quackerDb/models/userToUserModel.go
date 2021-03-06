package models

type UserToUser struct {
	UserId     int
	FollowerId int
}

func (utu UserToUser) IsModel() bool {
	return true
}
