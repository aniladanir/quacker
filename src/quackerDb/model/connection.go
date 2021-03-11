package model

type Connection struct {
	tableName  struct{} `pg:"connections"`
	Id         int      `pg:",pk"`
	FolloweeId int      `pg:"user_id,on_delete:RESTRICT"`
	FollowerId int      `pg:"on_delete:RESTRICT"`
}

func (c Connection) IsModel() bool {
	return true
}
