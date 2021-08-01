package controller

import (
	"fiber_news/models"
	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (u *UserController) Create(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if userRes, err := user.Create(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.JSON(userRes)
	}
}

func (u *UserController) LoadAll(c *fiber.Ctx) error {
	user := new(models.User)
	return c.JSON(user.LoadAll())
}

func (u *UserController) Load(c *fiber.Ctx) error {
	user := new(models.User)
	if resUser, err := user.Load(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.JSON(resUser)
	}
}

func (u *UserController) Delete(c *fiber.Ctx) error {
	var user models.User
	if msg, err := user.Delete(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.JSON(fiber.Map{"message": msg})
	}
}
