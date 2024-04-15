package controller

import (
	"github.com/gin-gonic/gin"

	constants "jaguar/trading/constants"
	service "jaguar/trading/service"
	"net/http"
)

func MountMarginController(router *gin.Engine) {
	controller := router.Group("/margin")
	marginService := service.NewMarginBotServiceImpl()
	{
		controller.GET("", getMargin)
		controller.GET("/getAll", getAllMargin)
		controller.POST("/publishMessage", marginService.PublishMessage)
	}
}

func getMargin(c *gin.Context) {
	c.JSON(http.StatusOK, "Margin found")
}

func getAllMargin(c *gin.Context) {
	c.JSON(http.StatusOK, constants.GetMarginData())
}
