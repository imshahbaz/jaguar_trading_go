package main

import (
	"log"
	"os"
	"zuma/src/configs/constants"
	"zuma/src/configs/threadpool"
	"zuma/src/controller"
	"zuma/src/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var zuma *fiber.App

func init() {
	defer handleDefer()
	zuma = fiber.New(fiber.Config{
		AppName: "Zuma",
	})

	zuma.Use(logger.New())
	threadpool.InitPool()
	controller.InitController(zuma)
	marginService := service.NewMarginService()
	constants.GsheetMargin = marginService.GetMarginData(constants.GsheetUrl)
}

func main() {
	defer handleDefer()
	if err := zuma.Listen(constants.PORT); err != nil {
		log.Panicf("Error starting server %v", err)
	}
}

func handleDefer() {
	if r := recover(); r != nil {
		if r == "" {
			return
		} else {
			log.Printf("Shutting down because of panic %v", r)
			shutdown()
		}
	}
}

func shutdown() {
	if zuma != nil {
		zuma.Shutdown()
	}
	os.Exit(0)
}
