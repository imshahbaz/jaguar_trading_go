package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"zuma/src/models/dto"
	"zuma/src/util"
)

type MarginService interface {
	GetMarginData(url string) map[string]dto.MarginDto
}

type marginServiceImpl struct{}

var marginServiceInstance *marginServiceImpl

func NewMarginService() MarginService {
	once.Do(func() {
		marginServiceInstance = &marginServiceImpl{}
	})
	return marginServiceInstance
}

func (service *marginServiceImpl) GetMarginData(url string) map[string]dto.MarginDto {

	resp := util.Get(url)
	if resp == nil {
		log.Panic("Error calling gsheet for margin")
	}

	file, err := os.Create("dataset.csv")
	if err != nil {
		log.Panicf("Error creating temp file %v", err)
	}

	defer func() {
		file.Close()
		err := os.Remove("dataset.csv")
		if err != nil {
			log.Panicf("failed to remove temporary file: %v", err)
		}
	}()

	_, err = file.Write(resp.Body())
	if err != nil {
		log.Panicf("failed to write response body to file: %v", err)
	}

	file.Seek(0, 0)

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Panicf("failed to read CSV data: %v", err)
	}

	var gSheetMargin map[string]dto.MarginDto = make(map[string]dto.MarginDto)
	for _, row := range rows[1:] {
		percent, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			continue
		}

		if percent >= 60 {
			symbol := row[1]
			gSheetMargin[symbol] = dto.MarginDto{
				Name:    row[0],
				Symbol:  symbol,
				Percent: percent,
			}
		}
	}

	return gSheetMargin
}
