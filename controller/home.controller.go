package controller

import (
	"fiber_news/models"

	"github.com/gofiber/fiber/v2"
)

type HomeController struct{}

func (u *HomeController) Home(c *fiber.Ctx) error {
	var news models.News
	return c.Render("index", fiber.Map{
		"Title":  "Hello Popo",
		"Newses": news.LoadAll(),
	}, "layouts/main")
}

func (u *HomeController) NewsPage(c *fiber.Ctx) error {
	var news models.News
	if news, err := news.Load(c.Params("id")); err != nil {
		return c.SendString("Go fuck urself")
	} else {
		return c.Render("news", fiber.Map{
			"Title": news.Title,
			"News":  news,
		})
	}
}
