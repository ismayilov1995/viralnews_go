package routers

import (
	"fiber_news/controller"
	"github.com/gofiber/fiber/v2"
)

var u = new(controller.UserController)

func UserRouter(r fiber.Router) {
	r.Get("/", u.LoadAll)
	r.Post("/signup", u.Create)
	r.Get("/:id", u.Load)
	r.Delete("/:id", u.Delete)
}
