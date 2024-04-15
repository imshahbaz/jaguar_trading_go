package service

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	constants "jaguar/trading/constants"

	"github.com/gin-gonic/gin"

	models "jaguar/trading/model"
)

const api = "https://apiway.ai/webhooks/catch/661b9ad15daec/webhooks-app"

type MarginBotService interface {
	PublishMessage(c *gin.Context)
}

type MarginBotServiceImpl struct{}

func NewMarginBotServiceImpl() *MarginBotServiceImpl {
	return &MarginBotServiceImpl{}
}

func (d *MarginBotServiceImpl) PublishMessage(c *gin.Context) {
	strategy := c.Query("strategy")
	var body models.ChartInk
	c.ShouldBindJSON(&body)
	data := constants.GetMarginData()

	list := strings.Split(body.Stocks, ",")

	message := ""
	for _, stock := range list {
		if value, ok := data[strings.TrimSpace(stock)]; ok {
			message += fmt.Sprintf("\n%s (%s) %d", value.Name, value.Symbol, value.Percent)
		}
	}

	formData := url.Values{}
	formData.Set("name", strategy)
	formData.Set("stocks", message)
	encodedData := formData.Encode()
	res, err := http.Post(api, "application/x-www-form-urlencoded", strings.NewReader(encodedData))
	fmt.Println(res)
	fmt.Println(err)
}
