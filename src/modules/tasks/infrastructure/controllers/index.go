package controllers

import (
	"github.com/Castelblanco/goland-clean/src/modules/common/middlewares"
	"github.com/Castelblanco/goland-clean/src/modules/common/responses/success"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/app/services"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/domain/entities"
	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	Services services.TaskServices
}

func (c TaskController) GetAll(ctx *fiber.Ctx) error {
	tasks, err := c.Services.GetAll()

	if err != nil {
		return middlewares.ErrorHandler(ctx, err)
	}

	return ctx.JSON(success.NewListResponse(
		tasks, len(tasks), fiber.StatusOK,
	))
}

func (c TaskController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	task, err := c.Services.GetById(id)

	if err != nil {
		return middlewares.ErrorHandler(ctx, err)
	}

	return ctx.JSON(success.ApiResponse{
		Item:   task,
		Status: fiber.StatusOK,
	})
}

func (c TaskController) CreateOne(ctx *fiber.Ctx) error {
	newTask := new(entities.TaskDOM)

	if err := ctx.BodyParser(newTask); err != nil {
		return err
	}

	task, err := c.Services.CreateOne(*newTask)

	if err != nil {
		return middlewares.ErrorHandler(ctx, err)
	}

	return ctx.JSON(success.ApiResponse{
		Item:   task,
		Status: fiber.StatusOK,
	})
}
