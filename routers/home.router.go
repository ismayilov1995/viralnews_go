package routers

import (
	"fiber_news/controller"

	"github.com/gofiber/fiber/v2"
)

var h = new(controller.HomeController)

func HomeRouter(r fiber.Router) {
	r.Get("/", h.Home)
	r.Get("/news/:id", h.NewsPage)
}
