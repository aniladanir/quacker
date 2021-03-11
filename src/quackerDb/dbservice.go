package quackerDb

import (
	"errors"
	"fmt"

	"github.com/aniladanir/quacker/quackerDb/model"
	"github.com/aniladanir/quacker/quackerDb/repo"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DbService struct {
	UserRepo  repo.Repo
	QuackRepo repo.Repo
	LikeRepo  repo.Repo
	ConnRepo  repo.Repo
}

var dbService_Singleton *DbService
var database *pg.DB

func GetService() (*DbService, error) {

	if database == nil {
		return nil, errors.New("No connection to database")
	}

	if dbService_Singleton != nil {
		return dbService_Singleton, nil
	}

	dbService := DbService{
		UserRepo: &repo.UserRepo{
			Db: database,
		},
		QuackRepo: &repo.QuackRepo{
			Db: database,
		},
		LikeRepo: &repo.LikeRepo{
			Db: database,
		},
		ConnRepo: &repo.ConnRepo{
			Db: database,
		},
	}

	dbService_Singleton = &dbService

	return dbService_Singleton, nil

}

func Connect(opts *pg.Options) error {
	db := pg.Connect(opts)
	if db == nil {
		fmt.Println("Failed to connect to database")
		return errors.New("Failed to connect to database")
	}

	database = db

	return nil
}

func CreateSchema() error {

	orm.RegisterTable((*model.Connection)(nil))

	models := []interface{}{
		(*model.User)(nil),
		(*model.Quack)(nil),
		(*model.Like)(nil),
		(*model.Connection)(nil),
	}
	for _, model := range models {
		err := database.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        true,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func Close() error {
	dbService_Singleton = nil
	err := database.Close()

	if err != nil {
		return err
	}

	return nil
}
