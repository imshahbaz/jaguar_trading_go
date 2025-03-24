package controller

import (
	"sync"
	"time"
	"zuma/src/configs/threadpool"
	"zuma/src/service"

	"github.com/gofiber/fiber/v2"
)

var telegramControllerInstance *telegramControllerImpl
var once sync.Once

type TelegramController interface {
	PublishMessage(ctx *fiber.Ctx) error
}

type telegramControllerImpl struct {
	TelegramService service.TelegramService
}

func NewTelegramController(telegramService service.TelegramService) TelegramController {
	once.Do(func() {
		telegramControllerInstance = &telegramControllerImpl{
			TelegramService: telegramService,
		}
	})
	return telegramControllerInstance
}

func (service *telegramControllerImpl) PublishMessage(ctx *fiber.Ctx) error {
	durationToAdd := 5*time.Hour + 30*time.Minute
	indianNow := time.Now().Add(durationToAdd)

	indianDayString := indianNow.Weekday().String()

	if indianDayString == "Saturday" || indianDayString == "Sunday" {
		return ctx.Status(fiber.StatusOK).JSON("WeekEnd")
	}

	threadpool.Pool.Submit(func() {
		service.TelegramService.SendMessageV2()
	})

	return ctx.Status(fiber.StatusOK).JSON("Message Published Successfully")
}
