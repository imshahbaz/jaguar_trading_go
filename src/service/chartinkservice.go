package service

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"zuma/src/models/dto"

	"github.com/chromedp/chromedp"
)

type ChartInkService interface {
	GetChartinkData(scanClause string) ([]dto.ChartInkDto, error)
	GetChartinkHtmlData(url string) ([]string, error)
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

func (service *chartInkServiceImpl) GetChartinkHtmlData(url string) ([]string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var symbols []string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("table#DataTables_Table_0 tbody", chromedp.ByQuery),
		chromedp.Evaluate(`
			(function() {
				var symbols = [];
				document.querySelectorAll("table#DataTables_Table_0 tbody tr").forEach(function(row) {
					var link = row.cells[1].querySelector("a");
					if (link) {
						var href = link.href;
						var urlParams = new URLSearchParams(href.split('?')[1]);
						var symbol = urlParams.get('symbol');
						if (symbol) {
							symbols.push(symbol);
						}
					}
				});
				return symbols;
			})();
		`, &symbols),
	)

	if err != nil {
		log.Printf("Failed to load page or extract symbols: %v", err)
		return nil, err
	}

	return symbols, nil
}
