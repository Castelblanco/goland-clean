package routers

import (
	"github.com/Castelblanco/goland-clean/src/modules/tasks/app/services"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/controllers"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/storages/sql/implementation"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/storages/sql/wrappers"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/tools"
	"github.com/gofiber/fiber/v2"
)

var createdId = tools.CreatedId
var repository = implementation.TaskSqlRepository{
	Wrappers: wrappers.TaskWrappers{},
}

var dependencies = services.Dependencies{
	Repository: repository,
	CreatedId:  createdId,
}

var service = services.TaskServices{
	Dependencies: dependencies,
}

var controller = controllers.TaskController{
	Services: service,
}

func InitTasksControllers(v1 fiber.Router) {
	router := v1.Group("tasks")

	router.Get("/get-all", controller.GetAll)
	router.Get("/get-one/:id", controller.GetById)
	router.Post("/create-one", controller.CreateOne)
	// router.Put("/update-one/:id", controller.UpdateOne)
	// router.Delete("/delete-one/:id", controller.DeleteOne)
}
