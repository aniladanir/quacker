package repos

import (
	"github.com/aniladanir/quacker/quackerDb/model"
	"github.com/go-pg/pg/v10"
)

type QuackRepo struct {
	Db *pg.DB
}

func (r *QuackRepo) Add(m model.Model) error {
	_, err := r.Db.Model(m.(*model.Quack)).Insert()

	if err != nil {
		return err
	}

	return nil
}

func (r *QuackRepo) Get(id int) (model.Model, error) {
	quack := model.Quack{Id: id}
	err := r.Db.Model(&quack).WherePK().Select()

	if err != nil {
		return quack, err
	}

	return quack, nil
}

func (r *QuackRepo) Update(m model.Model) error {
	_, err := r.Db.Model(m.(*model.Quack)).UpdateNotZero()

	if err != nil {
		return err
	}

	return nil
}

func (r *QuackRepo) Remove(id int) error {
	quack := model.Quack{Id: id}

	_, err := r.Db.Model(&quack).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
