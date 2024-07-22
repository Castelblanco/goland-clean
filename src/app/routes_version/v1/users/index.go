package users

import (
	"github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/routers"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(v1 fiber.Router) {
	routers.InitUsersControllers(v1)
}
