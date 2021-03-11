package repo

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

func (r *UserRepo) GetByTag(tag string) (model.Model, error) {
	var user model.User
	if err := r.Db.Model(&user).Where("tag = ?", tag).Select(); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) GetFollowers(id int) ([]model.User, error) {
	var followers = make([]model.User, 1)
	_, err := r.Db.Query(&followers, "SELECT * FROM users WHERE id IN "+
		"(SELECT follower_id FROM connections WHERE user_id = ?)", id)
	if err != nil {
		return nil, err
	}

	return followers, nil
}

func (r *UserRepo) GetFollowings(id int) ([]model.User, error) {
	var followings = make([]model.User, 1)
	_, err := r.Db.Query(&followings, "SELECT * FROM users WHERE id IN "+
		"(SELECT user_id FROM connections WHERE follower_id = ?)", id)
	if err != nil {
		return nil, err
	}

	return followings, nil
}

func (r *UserRepo) GetLikedQuacks(id int) ([]model.Quack, error) {
	var quacks = make([]model.Quack, 1)
	_, err := r.Db.Query(&quacks, "SELECT * FROM quacks WHERE id IN "+
		"(SELECT quack_id FROM likes WHERE user_id = ?)", id)
	if err != nil {
		return nil, err
	}

	return quacks, nil
}

func (r *UserRepo) Update(id int, m model.Model) error {
	_, err := r.Db.Model(&model.User{Id: id}).WherePK().UpdateNotZero(m.(*model.User))

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Remove(id int) (model.Model, error) {
	user := model.User{Id: id}
	_, err := r.Db.Model(&user).WherePK().Delete()
	if err != nil {
		return nil, err
	}

	return user, nil
}
