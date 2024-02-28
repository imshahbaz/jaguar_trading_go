package main

import (
	"fmt"
	"net/http"

	router "jaguar/trading/routers"

	"os"

	"github.com/gin-gonic/gin"

	constants "jaguar/trading/constants"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	constants.Init()
}

func main() {
	port := os.Getenv("PORT")
	app := gin.Default()

	app.GET("/", hello)
	router.MountRouter(app)

	fmt.Println("Started Jaguar Trading")
	app.Run(":" + port)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from jaguar trading")
}
