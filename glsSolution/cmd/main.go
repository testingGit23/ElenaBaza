package main

import (
	"ElenaBaza/glsSolution/pkg/readingCSV"
	"fmt"
)

func main() {
	rows := readingCSV.ReadValutes("Valutes.csv")
	valutes := readingCSV.CurrenciesConverter(rows)
	for index := 0; index < len(valutes); index++ {
		fmt.Println(valutes[index])
	}

	rowsV := readingCSV.ReadValutes("Companies.csv")
	companies := readingCSV.CurrenciesPerCompanyConverter(rowsV)
	for index := 0; index < len(companies); index++ {
		fmt.Println(companies[index])
	}
}
