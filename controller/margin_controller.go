package controller

import (
	"github.com/gin-gonic/gin"

	constants "jaguar/trading/constants"
	"net/http"
)

func MountMarginController(router *gin.Engine) {
	controller := router.Group("/margin")
	{
		controller.GET("", getMargin)
		controller.GET("/getAll", getAllMargin)
	}
}

func getMargin(c *gin.Context) {
	c.JSON(http.StatusOK, "Margin found")
}

func getAllMargin(c *gin.Context) {
	c.JSON(http.StatusOK, constants.GetMarginData())
}
