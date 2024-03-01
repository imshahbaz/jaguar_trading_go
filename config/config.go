package config

import (
	"os"

	constants "jaguar/trading/constants"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from jaguar trading")
}

var Port string
var MongoClient *mongo.Client

func Init() {
	godotenv.Load()
	release_mode := os.Getenv("RELEASE_MODE")
	gin.SetMode(release_mode)
	Port = os.Getenv("PORT")

	constants.Init()
	MongoClient = InitMongoDB()
}
