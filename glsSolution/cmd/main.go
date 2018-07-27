package main

import (
	"ElenaBaza/glsSolution/pkg/calculate"
	"ElenaBaza/glsSolution/pkg/readingCSV"
	"ElenaBaza/glsSolution/pkg/writingTotals"
)

func main() {
	rowsFromCompanies := readingCSV.CurrenciesPerCompanyReader("asd/Companies.csv")
	rowsFromValutes := readingCSV.CurrenciesReader("asd/Valutes.csv")

	totalsPerCompanyMap := calculate.CalculatingPerCompani(rowsFromCompanies, rowsFromValutes)

	writingTotals.WriteCSV("asd/Companies.csv", totalsPerCompanyMap)

}
