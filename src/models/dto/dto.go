package dto

type ChartInkDto struct {
	NseCode string `json:"nsecode"`
	Name    string `json:"name"`
}

type MarginDto struct {
	Name    string
	Symbol  string
	Percent float64
}
