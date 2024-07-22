package routers

import (
	"github.com/Castelblanco/goland-clean/src/modules/users/app/services"
	"github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/controllers"
	"github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/storages/sql/implementation"
	"github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/storages/sql/wrapper"
	"github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/tools"
	"github.com/gofiber/fiber/v2"
)

var createdId = tools.CreatedId
var repository = implementation.UsersSqlRepository{
	Wrapper: wrapper.UsersWrapper{},
}

var dependencies = services.Dependencies{
	Repository: repository,
	CreatedId:  createdId,
}
var allservices = services.UsersServices{
	Dependencies: dependencies,
}

var controller = controllers.UsersControllers{
	Services: allservices,
}

func InitUsersControllers(v1 fiber.Router) {
	router := v1.Group("users")

	router.Get("/get-all", controller.GetAll)
	router.Get("/get-one/:id", controller.GetById)
	router.Post("/create-one", controller.CreateOne)
	router.Put("/update-one/:id", controller.UpdateOne)
	router.Delete("/delete-one/:id", controller.DeleteOne)
}
