package main

import (
	"fmt"
	"os"

	"github.com/aniladanir/quacker/quackerDb"
	"github.com/aniladanir/quacker/quackerDb/model"
	"github.com/aniladanir/quacker/quackerDb/repo"
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

	users, err := dbService.UserRepo.(*repo.UserRepo).GetFollowings(1)
	fmt.Println(users)

}

func InsertMockupData(dbService *quackerDb.DbService) {
	users := []model.User{
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
	quacks := []model.Quack{
		{
			Id:      1,
			Content: "QUAAAACK",
			UserId:  1,
		},
		{
			Id:      2,
			Content: "asd",
			UserId:  1,
		},
		{
			Id:      3,
			Content: "quack",
			UserId:  2,
		},
		{
			Id:      4,
			Content: "ADFGAD",
			UserId:  2,
		},
		{
			Id:      5,
			Content: "quack",
			UserId:  3,
		},
		{
			Id:      6,
			Content: "tut",
			UserId:  3,
		},
	}
	likes := []model.Like{
		{
			Id:      1,
			UserId:  1,
			QuackId: 3,
		},
		{
			Id:      2,
			UserId:  1,
			QuackId: 6,
		},
		{
			Id:      3,
			UserId:  2,
			QuackId: 1,
		},
		{
			Id:      4,
			UserId:  2,
			QuackId: 5,
		},
		{
			Id:      5,
			UserId:  3,
			QuackId: 2,
		},
		{
			Id:      6,
			UserId:  3,
			QuackId: 4,
		},
	}
	conns := []model.Connection{
		{
			Id:         1,
			FolloweeId: 1,
			FollowerId: 2,
		},
		{
			Id:         2,
			FolloweeId: 1,
			FollowerId: 3,
		},
		{
			Id:         3,
			FolloweeId: 2,
			FollowerId: 1,
		},
		{
			Id:         4,
			FolloweeId: 2,
			FollowerId: 3,
		},
		{
			Id:         5,
			FolloweeId: 3,
			FollowerId: 1,
		},
		{
			Id:         6,
			FolloweeId: 3,
			FollowerId: 2,
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

	for _, like := range likes {
		err := dbService.LikeRepo.Add(&like)
		if err != nil {
			panic(err)
		}
	}

	for _, conn := range conns {
		err := dbService.ConnRepo.Add(&conn)
		if err != nil {
			panic(err)
		}
	}
}
