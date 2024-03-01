package service

import (
	"context"
	"net/http"

	constant "jaguar/trading/constants"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DividendService interface {
	GetDividend(c *gin.Context)
	SaveDividend(c *gin.Context)
}

type DividendServiceImpl struct {
	mongoClient *mongo.Client
}

func NewDividendServiceImpl(client *mongo.Client) *DividendServiceImpl {
	return &DividendServiceImpl{mongoClient: client}
}

func (d *DividendServiceImpl) GetDividend(c *gin.Context) {
	var dividends []constant.Margin
	cursor, err := d.mongoClient.Database(os.Getenv("DB_NAME")).Collection("margin").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}
	cursor.All(context.TODO(), &dividends)
	c.JSON(http.StatusOK, dividends)
}

func (d *DividendServiceImpl) SaveDividend(c *gin.Context) {
	requestBody := constant.Margin{}
	err := c.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Request Body "+err.Error())
		return
	}
	result, err := d.mongoClient.Database(os.Getenv("DB_NAME")).Collection("margin").InsertOne(context.TODO(), requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error saving  dividend details"+err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}
