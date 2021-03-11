package repo

import (
	"errors"

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
	//Get quack
	var quack *model.Quack
	_, err := r.Db.Query(quack, "SELECT * FROM quacks WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	//Add user info
	_, err = r.Db.Query(&quack.UserTag, &quack.UserName,
		"SELECT tag, name FROM users WHERE id = ?", quack.UserId) //REVISIT use JOIN instead of two queries
	if err != nil {
		return nil, err
	}

	//Check if quack has parent
	if quack.ParentId == 0 {
		return quack, nil
	}

	hasParent := true

	var parent *model.Quack
	id = quack.ParentId
	parents = make([]model.Quack, 1)

	for hasParent {
		id = parent.ParentId
		_, err := r.Db.Query(parent, "SELECT * FROM quacks WHERE id = ?", id)
		if err != nil {
			return nil, err
		}

		if quack.ParentId == 0 {
			hasParent = false
		}
	}

	return quack, nil
}

func (r *QuackRepo) Update(id int, m model.Model) error {
	return errors.New("Quacks can't be updated")
}

func (r *QuackRepo) Remove(id int) (model.Model, error) {
	quack := model.Quack{Id: id}
	_, err := r.Db.Model(&quack).WherePK().Delete()
	if err != nil {
		return nil, err
	}

	return quack, nil
}
