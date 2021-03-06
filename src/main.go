package main

import (
	"os"

	"github.com/aniladanir/quacker/quackerDb"
	"github.com/aniladanir/quacker/quackerDb/models"
	"github.com/go-pg/pg/v10"
)

func main() {

	//Setup connection to database
	dbOptions := &pg.Options{
		User:     os.Args[1],
		Password: os.Args[2],
	}

	err := quackerDb.Connect(dbOptions)
	if err != nil {
		panic(err)
	}

	defer quackerDb.Close()

	//Create tables
	err = quackerDb.CreateSchema()
	if err != nil {
		panic(err)
	}

	//Get database service for crud operations
	dbService, err := quackerDb.GetService()
	if err != nil {
		panic(err)
	}

	InsertMockupData(dbService)

	/*
		apiConfig := fiber.Config{
			ServerHeader:  "Quacker",
			CaseSensitive: true,
			StrictRouting: true,
		}
		app := fiber.New(apiConfig)
	*/
}

func InsertMockupData(dbService *quackerDb.DbService) {
	users := []models.User{
		{
			Id:          1,
			Tag:         "father",
			Email:       "god@heaven.com",
			Password:    "1",
			Name:        "God",
			Description: "Creator of all",
		},
		{
			Id:          2,
			Tag:         "jesus",
			Email:       "jesus@earth.com",
			Password:    "2",
			Name:        "Jesus Christ",
			Description: "",
		},
		{
			Id:          3,
			Tag:         "muhammad",
			Email:       "muhammad@earth.com",
			Password:    "3",
			Name:        "Muhammad",
			Description: "",
		},
	}
	quacks := []models.Quack{
		{
			Id:      1,
			Content: "QUAAAACK",
			UserId:  1,
		},
		{
			Id:      2,
			Content: "quack",
			UserId:  2,
		},
		{
			Id:      3,
			Content: "quack",
			UserId:  3,
		},
	}

	for _, user := range users {
		err := dbService.UserRepo.Add(&user)
		if err != nil {
			panic(err)
		}
	}

	for _, quack := range quacks {
		err := dbService.QuackRepo.Add(&quack)
		if err != nil {
			panic(err)
		}
	}
}
