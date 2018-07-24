package main

import (
	"ElenaBaza/glsSolution/pkg/readingCSV"
	"fmt"
)

func main() {
	rows := readingCSV.ReadValutes("Valutes.csv")
	valutes := readingCSV.CurrenciesConverter(rows)
	fmt.Println(valutes)
	rowsV := readingCSV.ReadValutes("Companies.csv")
	companies := readingCSV.CurrenciesPerCompanyConverter(rowsV)
	fmt.Println(companies)
}
