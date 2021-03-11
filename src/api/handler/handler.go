package handler

import (
	"net/http"
	"strconv"

	"github.com/aniladanir/quacker/api/data"
	"github.com/aniladanir/quacker/quackerDb/repo"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

const (
	_invalidId   = "Id is in invalid form"
	_notFound    = "Not found"
	_invalidData = "Missing or incompatible data"
	_internal    = "Internal error"
	_created     = "Created successfully"
	_found       = "Found"
	_updated     = "Updated successfully"
	_deleted     = "Deleted successfully"
)

//User handlers
func AddUser(c *fiber.Ctx, r *repo.UserRepo) error {
	var user data.UserEntry
	//Decode request body
	if err := c.BodyParser(user); err != nil {
		c.Status(http.StatusBadRequest).JSON(&response{
			Entity:  "User",
			Message: _invalidData,
		})
		return err
	}

	//Add user to database
	if err := r.Add(user.ToModel()); err == nil {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _internal,
		})
	}

	return c.Status(http.StatusOK).JSON(&response{
		Entity:  "User",
		Message: _created,
	})

}

func UpdateUser(c *fiber.Ctx, r *repo.UserRepo) error {
	//Get id from http route
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _invalidId,
		})
		return err
	}

	//Decode request body
	var user data.UserEntry
	if err := c.BodyParser(user); err != nil {
		c.Status(http.StatusBadRequest).JSON(&response{
			Entity:  "User",
			Message: _invalidData,
		})
		return err
	}

	//Update user
	if err := r.Update(id, user.ToModel()); err == pg.ErrNoRows {
		c.Status(http.StatusInternalServerError).JSON((&response{
			Entity:  "User",
			Message: _notFound,
		}))
		return err
	} else if err != nil {
		c.Status(http.StatusInternalServerError).JSON((&response{
			Entity:  "User",
			Message: _internal,
		}))
		return err
	}

	return c.Status(http.StatusOK).JSON(&response{
		Entity:  "User",
		Message: _updated,
	})
}

func DeleteUser(c *fiber.Ctx, r *repo.UserRepo) error {
	//Get id from http route
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&response{
			Entity:  "User",
			Message: _invalidId,
		})
		return err
	}

	//Remove user
	user, err := r.Remove(id)
	if err == pg.ErrNoRows {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _notFound,
		})
		return err
	} else if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _internal,
		})
		return err
	}

	return c.Status(http.StatusOK).JSON(&response{
		Entity:  "User",
		Message: _deleted,
		Data:    user,
	})
}

func GetUserById(c *fiber.Ctx, r *repo.UserRepo) error {
	//Get id from http route
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&response{
			Entity:  "User",
			Message: _invalidId,
		})
		return err
	}

	//Get user
	user, err := r.Get(id)
	if err == pg.ErrNoRows {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _notFound,
		})
		return err
	} else if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _internal,
		})
		return err
	}

	return c.Status(http.StatusOK).JSON(&response{
		Entity:  "User",
		Message: _found,
		Data:    user,
	})

}

func GetUserByTag(c *fiber.Ctx, r *repo.UserRepo) error {
	tag := c.Params("tag")

	user, err := r.GetByTag(tag)
	if err == pg.ErrNoRows {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _notFound,
		})
		return err
	} else if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _internal,
		})
		return err
	}

	return c.Status(http.StatusOK).JSON(&response{
		Entity:  "User",
		Message: _found,
		Data:    user,
	})
}

func GetTimeline(c *fiber.Ctx, userR *repo.UserRepo) {
	/*
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(&response{
				Status:  _error,
				Message: "Id is in invalid form",
			})
		}
		user, err := userR.Get(id)
		user.(*model.User).Followers[1].
	*/
}

func GetFollowersById(c *fiber.Ctx, r *repo.UserRepo) error {
	//Get id from http route
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&response{
			Entity:  "User",
			Message: _invalidId,
		})
		return err
	}

	//Get followers
	followers, err := r.GetFollowers(id)
	if err == pg.ErrNoRows {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _notFound,
		})
		return err
	} else if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _internal,
		})
		return err
	}

	return c.Status(http.StatusOK).JSON(&response{
		Entity:  "User",
		Message: _found,
		Data:    followers,
	})

}

func GetFollowingsById(c *fiber.Ctx, r *repo.UserRepo) error {
	//Get id from http route
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&response{
			Entity:  "User",
			Message: _invalidId,
		})
		return err
	}

	//Get followers
	followings, err := r.GetFollowings(id)
	if err == pg.ErrNoRows {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _notFound,
		})
		return err
	} else if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _internal,
		})
		return err
	}

	return c.Status(http.StatusOK).JSON(&response{
		Entity:  "User",
		Message: _found,
		Data:    followings,
	})
}

func GetLikedQuacksByUserId(c *fiber.Ctx, r *repo.UserRepo) error {
	//Get id from http route
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&response{
			Entity:  "User",
			Message: _invalidId,
		})
		return err
	}

	//Get liked quacks
	quacks, err := r.GetLikedQuacks(id)
	if err == pg.ErrNoRows {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _notFound,
		})
		return err
	} else if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&response{
			Entity:  "User",
			Message: _internal,
		})
		return err
	}

	return c.Status(http.StatusOK).JSON(&response{
		Entity:  "Quack",
		Message: _found,
		Data:    quacks,
	})
}

func GetWithRepliesByUserTag(c *fiber.Ctx, r *repo.UserRepo) {
	//Get tag from http route
	tag := c.Params("tag") //REVISIT Check tag for unallowed characters (in middleware?)
	r.GetLikedQuacks()

}

//Quack handlers
func AddQuack(c *fiber.Ctx, r *repo.QuackRepo) {

}

func UpdateQuack(c *fiber.Ctx, r *repo.QuackRepo) {

}

func DeleteQuack(c *fiber.Ctx, r *repo.QuackRepo) {

}

func GetQuackById(c *fiber.Ctx, r *repo.QuackRepo) {

}

func GetRequacksByQuackId(c *fiber.Ctx, r *repo.QuackRepo) {

}

func GetLikesByQuackId(c *fiber.Ctx, r *repo.QuackRepo) {

}

func GetQuotedByQuackId(c *fiber.Ctx, r *repo.QuackRepo) {

}

//Like handlers
func AddLike(c *fiber.Ctx) {

}

func DeleteLike(c *fiber.Ctx) {

}

//Connection handlers
func AddConnection(c *fiber.Ctx) {

}

func DeleteConnection(c *fiber.Ctx) {

}
