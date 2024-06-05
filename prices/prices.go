package prices

import (
	"fmt"
	"strconv"

	"github.com/qsmsoft/price-calc/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	taxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Canverting price to float failed!")
			fmt.Println(err)
			return
		}

		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices

}

func (job TaxIncludedPriceJob) Proccess() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.taxIncludedPrices = result

	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
