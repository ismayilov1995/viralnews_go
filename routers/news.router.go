package routers

import (
	"fiber_news/controller"

	"github.com/gofiber/fiber/v2"
)

var n = new(controller.NewsController)

func NewsRouter(r fiber.Router) {
	r.Get("/", n.LoadAll)
	r.Post("/create", n.Create)
	r.Get("/seed", n.Seed)
	r.Get("/reset", n.Reset)
	r.Get("/:id", n.Load)
	r.Delete("/:id", n.Delete)
}
