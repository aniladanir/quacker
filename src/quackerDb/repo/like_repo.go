package repo

import (
	"errors"

	"github.com/aniladanir/quacker/quackerDb/model"
	"github.com/go-pg/pg/v10"
)

type LikeRepo struct {
	Db *pg.DB
}

func (r *LikeRepo) Add(m model.Model) error {
	_, err := r.Db.Model(m.(*model.Like)).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (r *LikeRepo) Get(id int) (model.Model, error) {
	like := model.Like{Id: id}
	err := r.Db.Model(&like).WherePK().Select()

	if err != nil {
		return like, err
	}

	return like, nil
}

// Like has no update functionality for now
func (r *LikeRepo) Update(id int, m model.Model) error {
	return errors.New("Likes can't be updated")
}

func (r *LikeRepo) Remove(id int) (model.Model, error) {
	like := model.Like{Id: id}
	_, err := r.Db.Model(&like).WherePK().Delete()
	if err != nil {
		return nil, err
	}

	return like, nil
}
