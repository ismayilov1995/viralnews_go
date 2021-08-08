package routers

import (
	"fiber_news/controller"

	"github.com/gofiber/fiber/v2"
)

var h = new(controller.HomeController)

func HomeRouter(r fiber.Router) {
	r.Get("/", h.Home)
	r.Get("/news/create", h.NewsCreate)
	r.Post("/news/create", h.NewsCreate)
	r.Get("/news/:id", h.NewsPage)
}
