package repos

import (
	"github.com/aniladanir/quacker/quackerDb/model"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	Db *pg.DB
}

func (r *UserRepo) Add(m model.Model) error {
	_, err := r.Db.Model(m.(*model.User)).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Get(id int) (model.Model, error) {
	user := model.User{Id: id}
	err := r.Db.Model(&user).WherePK().Select()

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) Update(m model.Model) error {
	_, err := r.Db.Model(m.(*model.User)).WherePK().UpdateNotZero()

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Remove(id int) error {
	user := model.User{
		Id: id,
	}

	_, err := r.Db.Model(&user).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
