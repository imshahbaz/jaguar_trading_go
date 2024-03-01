package main

import (
	"fmt"

	router "jaguar/trading/routers"

	"github.com/gin-gonic/gin"

	config "jaguar/trading/config"
)

func init() {
	config.Init()
}

func main() {
	app := gin.Default()

	app.GET("/", config.Hello)
	router.MountRouter(app)

	fmt.Println("Started Jaguar Trading")
	app.Run(":" + config.Port)
}
