package main

import (
	"fmt"

	"github.com/qsmsoft/price-calc/cmdmanager"
	"github.com/qsmsoft/price-calc/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Proccess()
		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}

}
