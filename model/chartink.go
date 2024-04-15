package model

type ChartInk struct {
	Stocks string `json:"stocks" binding:"required"`
}
