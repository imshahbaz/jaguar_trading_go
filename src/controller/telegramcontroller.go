package controller

import (
	"log"
	"sync"
	"time"
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
	utcNow := time.Now().UTC()

	indianTimezone, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Println("Error loading time zone:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON("Error fetching time")
	}

	indianNow := utcNow.In(indianTimezone)

	indianDayString := indianNow.Weekday().String()

	if indianDayString == "Saturday" || indianDayString == "Sunday" {
		return ctx.Status(fiber.StatusOK).JSON("WeekEnd")
	}

	service.TelegramService.SendMessage()
	return ctx.Status(fiber.StatusOK).JSON("Message Published Successfully")
}
