package tasks

import (
	"github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/routers"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(v1 fiber.Router) {
	routers.InitTasksControllers(v1)
}
