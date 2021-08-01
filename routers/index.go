package routers

import "github.com/gofiber/fiber/v2"

func ApiRouters(r fiber.Router) {
	api := r.Group("/api/v1")
	UserRouter(api.Group("/user"))
	NewsRouter(api.Group("/news"))

}