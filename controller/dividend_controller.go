package controller

import (
	service "jaguar/trading/service"

	config "jaguar/trading/config"

	"github.com/gin-gonic/gin"
)

func MountDividendController(router *gin.Engine) {
	controller := router.Group("/dividend")
	dividendService := service.NewDividendServiceImpl(config.MongoClient)
	{
		controller.GET("/getAll", dividendService.GetDividend)
		controller.POST("/save", dividendService.SaveDividend)
	}
}
