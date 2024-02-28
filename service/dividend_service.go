package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DividendService interface {
	GetDividend(c *gin.Context)
}

type DividendServiceImpl struct {
}

func NewDividendServiceImpl() *DividendServiceImpl {
	return &DividendServiceImpl{}
}

func (d *DividendServiceImpl) GetDividend(c *gin.Context) {
	c.JSON(http.StatusOK, "dividends")
}
