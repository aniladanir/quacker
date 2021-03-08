package model

type Connection struct {
	tableName  struct{} `pg:"user_to_users"`
	FolloweeId int      `pg:"user_id,on_delete:RESTRICT"`
	FollowerId int      `pg:"on_delete:RESTRICT"`
}

func (c Connection) IsModel() bool {
	return true
}
