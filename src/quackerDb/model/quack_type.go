package model

type QuackType struct {
	tableName struct{} `pg:"quack_types"`
	id        int      `pg:",pk"`
	_type     string   `pg:"type"`
}

//Quack
//	Reply
//	Requack
//		Quoted quack
