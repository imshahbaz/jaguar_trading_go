package constants

import (
	"zuma/src/models/dto"
)

var PORT = ":8080"

const (
	GsheetUrl  = "https://docs.google.com/spreadsheets/d/e/2PACX-1vTb4OSG_m0d7SoEaVL70BwiO0VgwHKVLIXNOClkkJkXefRp33tYVOUAU_DXfwuLmFfJ-PmRI_qfIsHW/pub?output=csv"
	MessageUrl = "https://apiway.ai/webhooks/catch/661b9ad15daec/webhooks-app"
)

type strategy struct {
	Name   string
	Clause string
	Active bool
}

var Strategies = []strategy{
	{Name: "Vollinger", Clause: "( {46553} ( latest close > latest open and latest close > latest sma( latest close , 20 ) and latest volume > 1 day ago volume and latest low <= latest sma( latest close , 20 ) and latest close >= ( ( latest high - latest low ) * 0.7 ) + latest low ) )", Active: true},
}

var GsheetMargin = make(map[string]dto.MarginDto)

const (
	TelegramGroup   = "/telegram"
	TelegramPublish = "/publish"
)
