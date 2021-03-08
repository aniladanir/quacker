package repos

import (
	"github.com/aniladanir/quacker/quackerDb/models"
	"github.com/go-pg/pg/v10"
)

type QuackRepo struct {
	Db *pg.DB
}

func (r *QuackRepo) Add(m models.Model) error {
	_, err := r.Db.Model(m.(*models.Quack)).Insert()

	if err != nil {
		return err
	}

	return nil
}

func (r *QuackRepo) Get(id int) (models.Model, error) {
	quack := models.Quack{Id: id}
	err := r.Db.Model(&quack).WherePK().Select()

	if err != nil {
		return quack, err
	}

	return quack, nil
}

func (r *QuackRepo) Update(m models.Model) error {
	_, err := r.Db.Model(m.(*models.Quack)).UpdateNotZero()

	if err != nil {
		return err
	}

	return nil
}

func (r *QuackRepo) Remove(id int) error {
	quack := models.Quack{Id: id}

	_, err := r.Db.Model(&quack).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
