package repos

import (
	"github.com/aniladanir/quacker/quackerDb/models"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	Db *pg.DB
}

func (r *UserRepo) Add(m models.Model) error {
	_, err := r.Db.Model(m.(*models.User)).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Get(id int) (models.Model, error) {
	user := models.User{Id: id}
	err := r.Db.Model(&user).WherePK().Select()

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) Update(m models.Model) error {
	_, err := r.Db.Model(m.(*models.User)).WherePK().UpdateNotZero()

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Remove(id int) error {
	user := models.User{
		Id: id,
	}

	_, err := r.Db.Model(&user).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
