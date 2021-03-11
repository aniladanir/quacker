package api

import (
	"github.com/aniladanir/quacker/quackerDb"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, dbService *quackerDb.DbService) error {
	api := app.Group("/api")

	userRoute := api.Group("/user")
	userRoute.Post("/")
	userRoute.Put("/:id")
	userRoute.Delete("/:id")
	userRoute.Get("/id::id")
	userRoute.Get("/timeline/id::id")
	userRoute.Get("/followers/id::id")
	userRoute.Get("/followings:/id::id")
	userRoute.Get("/quacks/id::id")
	userRoute.Get("/withreplies/id::id")
	userRoute.Get("/likes/id::id")
	userRoute.Get("/tag::tag")
	userRoute.Get("/quacks/tag::tag")
	userRoute.Get("/withreplies/tag::tag")

	quackRoute := api.Group("/quack")
	quackRoute.Post("/")
	quackRoute.Put("/:id")
	quackRoute.Delete("/:id")
	quackRoute.Get("/id::id")
	quackRoute.Get("/replies/id::id")
	quackRoute.Get("/requacks/id::id")
	quackRoute.Get("/likes/id::id")
	quackRoute.Get("/quoted/id::id")

	connRoute := api.Group("connection")
	connRoute.Post("/")
	connRoute.Delete("/:id")

	likeRoute := api.Group("/like")
	likeRoute.Post("/")
	likeRoute.Delete("/:id")

	return nil
}
