package server

import (
	routesversion "github.com/Castelblanco/goland-clean/src/app/routes_version"
	v1 "github.com/Castelblanco/goland-clean/src/app/routes_version/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var App = fiber.New()

func middleware() {
	App.Use(logger.New())
	App.Use(requestid.New())
	App.Use(cors.New())
}

func routes() {
	routesversion.InitHealthyControllers(App)
	v1.InitRouterV1(App)
}

func Server() {
	middleware()
	routes()
	App.Listen(":5000")
}
