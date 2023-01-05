package routes

import (
	"github.com/francis560/api-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RandUsers(r fiber.Router) {

	r.Get("/users/:quantity", controllers.RandUsersHandler)

}
