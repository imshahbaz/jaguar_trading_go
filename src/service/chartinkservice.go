package service

import (
	"encoding/json"
	"os/exec"
	"zuma/src/models/dto"
)

type ChartInkService interface {
	GetChartinkData(scanClause string) ([]dto.ChartInkDto, error)
}

type chartInkServiceImpl struct{}

var chartInkServiceInstance *chartInkServiceImpl

func NewChartInkService() ChartInkService {
	once.Do(func() {
		chartInkServiceInstance = &chartInkServiceImpl{}
	})
	return chartInkServiceInstance
}

func (service *chartInkServiceImpl) GetChartinkData(scanClause string) ([]dto.ChartInkDto, error) {

	cmd := exec.Command("python", "./get_data.py", scanClause)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var result []dto.ChartInkDto
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
