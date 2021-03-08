package handlers

import (
	"github.com/aniladanir/quacker/db/repos"
	"github.com/gofiber/fiber/v2"
)

type QuackHandler struct {
	Repository repos.Repo
}

func (h *QuackHandler) GetAll(c *fiber.Ctx) error {
	return nil
}

func (h *QuackHandler) Get(c *fiber.Ctx) error {
	return nil
}

func (h *QuackHandler) Create(c *fiber.Ctx) error {
	return nil
}

func (h *QuackHandler) Update(c *fiber.Ctx) error {
	return nil
}

func (h *QuackHandler) Delete(c *fiber.Ctx) error {
	return nil
}
