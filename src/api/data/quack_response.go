package data

import "time"

//Use for get operation
type QuackResponse struct {
	Id           int
	QuackType    int
	Content      string
	DateCreated  time.Time
	OwnerTag     string
	OwnerName    string
	ReplyCount   int
	RequackCount int
	LikeCount    int
	Parents      []QuackResponse
	//Parents are in hierarchical order (first one in the array -> first parent etc.)
}
