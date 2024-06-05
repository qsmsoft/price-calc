package main

import (
	"fmt"

	"github.com/qsmsoft/price-calc/filemanager"
	"github.com/qsmsoft/price-calc/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(*fm, taxRate)
		priceJob.Proccess()
	}

}
