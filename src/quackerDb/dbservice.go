package quackerDb

import (
	"errors"
	"fmt"

	"github.com/aniladanir/quacker/quackerDb/models"
	"github.com/aniladanir/quacker/quackerDb/repos"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DbService struct {
	UserRepo  repos.Repo
	QuackRepo repos.Repo
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
		UserRepo: &repos.UserRepo{
			Db: database,
		},
		QuackRepo: &repos.QuackRepo{
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

	orm.RegisterTable((*models.UserToUser)(nil))

	models := []interface{}{
		(*models.User)(nil),
		(*models.Quack)(nil),
		(*models.Favorite)(nil),
		(*models.UserToUser)(nil),
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
