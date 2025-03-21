package controller

import (
	"zuma/src/configs/constants"
	"zuma/src/service"

	"github.com/gofiber/fiber/v2"
)

func InitController(app *fiber.App) {
	app.Get("/", defaultFunc)
	mapTelegramController(app)
}

func defaultFunc(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON("Welcome to ZUMA")
}

var telegramService = service.NewTelegramService()

func mapTelegramController(app *fiber.App) {
	controller := NewTelegramController(telegramService)
	telegram := app.Group(constants.TelegramGroup)
	telegram.Get(constants.TelegramPublish, controller.PublishMessage)
}
