package service

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"zuma/src/configs/constants"
	"zuma/src/models/dto"
	"zuma/src/util"
)

type TelegramService interface {
	SendMessage()
	SendMessageV2()
}

type telegramServiceImpl struct {
}

var telegramServiceInstance *telegramServiceImpl
var once sync.Once

func NewTelegramService() TelegramService {
	once.Do(func() {
		telegramServiceInstance = &telegramServiceImpl{}
	})
	return telegramServiceInstance
}

func (service *telegramServiceImpl) SendMessage() {
	chartInkService := NewChartInkService()
	for _, strategy := range constants.Strategies {
		if !strategy.Active {
			continue
		}

		data, err := chartInkService.GetChartinkData(strategy.Clause)
		if err != nil {
			log.Printf("Error fetching chartink data : %v", err)
			continue
		}

		if len(data) > 0 {
			var finallist []dto.MarginDto
			for _, stock := range data {
				s := constants.GsheetMargin[stock.NseCode]
				if s.Name == "" {
					continue
				}
				finallist = append(finallist, constants.GsheetMargin[stock.NseCode])
			}

			if len(finallist) > 0 {
				sort.Slice(finallist, func(i, j int) bool {
					return finallist[i].Percent > finallist[j].Percent
				})
				log.Println(finallist)
				message := "Strategy : " + strategy.Name + "\n"
				for _, margin := range finallist {
					message += fmt.Sprintf("\n%s ( %s ) %v\n", margin.Name, margin.Symbol, margin.Percent)
				}
				if message != "" {
					formData := map[string]string{
						"stocks": message,
					}
					util.Post(constants.MessageUrl, formData)
				}
			}
		}
	}
}

func (service *telegramServiceImpl) SendMessageV2() {
	chartInkService := NewChartInkService()
	for _, strategy := range constants.ChartInkScanners {
		if !strategy.Active {
			continue
		}
		result, err := chartInkService.GetChartinkHtmlData(strategy.Clause)
		if err != nil {
			continue
		}
		if len(result) > 0 {
			var finallist []dto.MarginDto
			for _, stock := range result {
				s := constants.GsheetMargin[stock]
				if s.Name == "" {
					continue
				}
				finallist = append(finallist, constants.GsheetMargin[stock])
			}

			if len(finallist) > 0 {
				sort.Slice(finallist, func(i, j int) bool {
					return finallist[i].Percent > finallist[j].Percent
				})
				log.Println(finallist)
				message := "Strategy : " + strategy.Name + "\n"
				for _, margin := range finallist {
					message += fmt.Sprintf("\n%s ( %s ) %v\n", margin.Name, margin.Symbol, margin.Percent)
				}
				if message != "" {
					formData := map[string]string{
						"stocks": message,
					}
					util.Post(constants.MessageUrl, formData)
				}
			}
		}
	}
}
