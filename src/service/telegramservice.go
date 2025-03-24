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
}

type telegramServiceImpl struct{}

var telegramServiceInstance *telegramServiceImpl
var once sync.Once

func NewTelegramService() TelegramService {
	once.Do(func() {
		telegramServiceInstance = &telegramServiceImpl{}
	})
	return telegramServiceInstance
}

func (service *telegramServiceImpl) SendMessage() {
	chartink := NewChartInkService()
	for _, strategy := range constants.Strategies {
		if !strategy.Active {
			continue
		}

		data, err := chartink.GetChartinkData(strategy.Clause)
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
