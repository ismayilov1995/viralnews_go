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
		return c.SendString("go fuck urself")
	} else {
		return c.Render("news", fiber.Map{
			"Title": news.Title,
			"News":  news,
		}, "layouts/main")
	}
}

func (u *HomeController) NewsCreate(c *fiber.Ctx) error {
	var author *models.User
	author, _ = author.Load(1)

	if c.Method() == "GET" {
		return c.Render("news_create", fiber.Map{
			"Title":  "Create News",
			"Author": author,
		}, "layouts/main")
	} else {
		var news models.News
		c.BodyParser(&news)
		news.AuthorID = 1
		if _, err := news.Create(); err != nil {
			return c.SendString(err.Error())
		} else {
			return c.JSON(news)
		}
	}
}
