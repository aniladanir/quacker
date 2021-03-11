package repo

import (
	"errors"

	"github.com/aniladanir/quacker/quackerDb/model"
	"github.com/go-pg/pg/v10"
)

type ConnRepo struct {
	Db *pg.DB
}

func (r *ConnRepo) Add(m model.Model) error {
	_, err := r.Db.Model(m.(*model.Connection)).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (r *ConnRepo) Get(id int) (model.Model, error) {
	conn := model.Connection{Id: id}
	err := r.Db.Model(&conn).WherePK().Select()

	if err != nil {
		return conn, err
	}

	return conn, nil
}

// Connection has no update functionality for now
func (r *ConnRepo) Update(id int, m model.Model) error {
	return errors.New("Connections can't be updated")
}

func (r *ConnRepo) Remove(id int) (model.Model, error) {
	conn := model.Connection{Id: id}
	_, err := r.Db.Model(&conn).WherePK().Delete()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
