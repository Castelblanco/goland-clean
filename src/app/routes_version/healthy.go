package routesversion

import (
	"github.com/Castelblanco/goland-clean/src/modules/common/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitHealthyControllers(app *fiber.App) {
	controllers := controllers.HealthyControllers{}

	controllers.Init()
	app.Get("/", controllers.Get)
	app.Get("/healthy", controllers.Healthy)
	app.Get("/readiness", controllers.Readiness)
}
