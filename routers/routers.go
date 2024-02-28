package routers

import (
	controllers "jaguar/trading/controller"

	"github.com/gin-gonic/gin"
)

func MountRouter(router *gin.Engine) {
	controllers.MountMarginController(router)
	controllers.MountDividendController(router)
}
