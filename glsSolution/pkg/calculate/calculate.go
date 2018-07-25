package calculate

import (
	"ElenaBaza/glsSolution/pkg/readingCSV"
	"strconv"
	"strings"
)

func CalculatingPerCompani() map[string]string {
	rows := readingCSV.ReadValutes("asd/Valutes.csv")
	valutes := readingCSV.CurrenciesConverter(rows)
	rowsV := readingCSV.ReadValutes("asd/Companies.csv")
	companies := readingCSV.CurrenciesPerCompanyConverter(rowsV)
	ret := make(map[string]string)
	for i := 0; i < len(companies); i++ {
		currenciesPerCompany := companies[i].CurrenciesForThisCompany
		var totalPerCompany float64
		for j := 0; j < len(currenciesPerCompany); j++ {
			currrentCurrency := currenciesPerCompany[j]
			var inDenars float64
			for k := 0; k < len(valutes); k++ {
				if strings.EqualFold(currrentCurrency.Currency, valutes[k].Currency) {
					inDenars = (currrentCurrency.InDenars * valutes[k].InDenars)
				} else {
					continue
				}
			}
			totalPerCompany += inDenars
		}
		ret[companies[i].CompanyName] = strconv.FormatFloat(totalPerCompany, 'f', -1, 64)
	}
	return ret
}
