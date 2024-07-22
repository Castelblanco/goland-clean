package controllers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Healthy struct {
	Message     string `json:"message"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

var healthy Healthy

type HealthyControllers struct{}

func (h HealthyControllers) Get(c *fiber.Ctx) error {
	return c.JSON(healthy)
}

func (h HealthyControllers) Healthy(c *fiber.Ctx) error {
	return c.JSON(healthy)
}

func (h HealthyControllers) Readiness(c *fiber.Ctx) error {
	return c.JSON(healthy)
}

func (h HealthyControllers) Init() {
	healthy = Healthy{
		Message:     fmt.Sprintf("%s running 👩‍💻", os.Getenv("APP_ID")),
		Version:     "0.0.1",
		Environment: os.Getenv("GO_ENV"),
	}
}
