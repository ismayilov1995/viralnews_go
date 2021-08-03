package controller

import (
	"fiber_news/models"

	"github.com/gofiber/fiber/v2"
)

type NewsController struct{}

func (n *NewsController) Create(c *fiber.Ctx) error {
	var news models.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if newsRes, err := news.Create(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.JSON(newsRes)
	}
}

func (n *NewsController) LoadAll(c *fiber.Ctx) error {
	news := new(models.News)
	return c.JSON(news.LoadAll())
}

func (n *NewsController) Load(c *fiber.Ctx) error {
	news := new(models.News)
	if newses, err := news.Load(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		return c.JSON(newses)
	}
}

func (u *NewsController) Delete(c *fiber.Ctx) error {
	var news models.News
	if msg, err := news.Delete(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.JSON(fiber.Map{"message": msg})
	}
}

func (n *NewsController) Seed(c *fiber.Ctx) error {
	var news models.News
	return c.JSON(news.Seed())
}

func (n *NewsController) Reset(c *fiber.Ctx) error {
	var news models.News
	news.Reset()
	return c.JSON(fiber.Map{"message": "removed"})
}
