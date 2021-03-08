package handlers

import (
	"github.com/aniladanir/quacker/db/repos"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Repository repos.Repo
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	return nil
}
