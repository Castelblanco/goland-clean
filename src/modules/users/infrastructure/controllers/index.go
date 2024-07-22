package controllers

import (
	"github.com/Castelblanco/goland-clean/src/modules/common/middlewares"
	"github.com/Castelblanco/goland-clean/src/modules/common/responses/success"
	"github.com/Castelblanco/goland-clean/src/modules/users/app/services"
	"github.com/Castelblanco/goland-clean/src/modules/users/domain/entities"
	"github.com/gofiber/fiber/v2"
)

type UsersControllers struct {
	Services services.UsersServices
}

func (c UsersControllers) GetAll(ctx *fiber.Ctx) error {
	users, err := c.Services.GetAll()
	if err != nil {
		return middlewares.ErrorHandler(ctx, err)
	}

	return ctx.JSON(success.NewListResponse(
		users, len(users), fiber.StatusOK,
	))
}

func (c UsersControllers) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.Services.GetById(id)

	if err != nil {
		return middlewares.ErrorHandler(ctx, err)
	}

	return ctx.JSON(success.ApiResponse{
		Item:   user,
		Status: fiber.StatusOK,
	})
}

func (c UsersControllers) CreateOne(ctx *fiber.Ctx) error {
	newUser := new(entities.UserDOM)

	if err := ctx.BodyParser(newUser); err != nil {
		return err
	}

	user, err := c.Services.CreateOne(*newUser)

	if err != nil {
		return middlewares.ErrorHandler(ctx, err)
	}

	return ctx.JSON(success.ApiResponse{
		Item:   user,
		Status: fiber.StatusCreated,
	})
}

func (c UsersControllers) UpdateOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	updateUser := new(entities.UserDOM)

	if err := ctx.BodyParser(updateUser); err != nil {
		return err
	}
	updateUser.Id = id
	user, err := c.Services.UpdateOne(*updateUser)

	if err != nil {
		return middlewares.ErrorHandler(ctx, err)
	}

	return ctx.JSON(success.ApiResponse{
		Item:   user,
		Status: fiber.StatusOK,
	})
}

func (c UsersControllers) DeleteOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.Services.DeleteOne(id)

	if err != nil {
		return middlewares.ErrorHandler(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
