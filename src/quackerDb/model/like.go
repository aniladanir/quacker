package model

type Like struct {
	tableName struct{} `pg:"likes"`
	Id        int      `pg:",pk"`
	UserId    int      `pg:"on_delete:RESTRICT"`
	QuackId   int      `pg:"on_delete:RESTRICT"`
}

func (f Like) IsModel() bool {

	return true
}
