package controller

import (
	service "jaguar/trading/service"

	"github.com/gin-gonic/gin"
)

func MountDividendController(router *gin.Engine) {
	controller := router.Group("/dividend")
	dividendService := service.NewDividendServiceImpl()
	{
		controller.GET("/getAll", dividendService.GetDividend)
	}
}
