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

var port string

func init() {
	godotenv.Load()
	constants.Init()
	config()
}

func main() {
	app := gin.Default()

	app.GET("/", hello)
	router.MountRouter(app)

	fmt.Println("Started Jaguar Trading")
	app.Run(":" + port)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from jaguar trading")
}

func config() {
	release_mode := os.Getenv("RELEASE_MODE")
	gin.SetMode(release_mode)
	port = os.Getenv("PORT")
}
