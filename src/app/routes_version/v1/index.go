package v1

import (
	"github.com/Castelblanco/goland-clean/src/app/routes_version/v1/tasks"
	"github.com/Castelblanco/goland-clean/src/app/routes_version/v1/users"
	"github.com/gofiber/fiber/v2"
)

func InitRouterV1(app *fiber.App) {
	router := app.Group("v1")

	users.InitRouter(router)
	tasks.InitRouter(router)
}
