package calculate

import (
	"ElenaBaza/glsSolution/pkg/readingCSV"
	"strconv"
	"strings"
)

func CalculatingPerCompani(companies []readingCSV.CurrenciesPerCompany, valutes []readingCSV.Currencies) map[string]string {
	ret := make(map[string]string)
	for i := 0; i < len(companies); i++ {
		currenciesPerCompany := companies[i].CurrenciesForThisCompany
		var totalPerCompany float64
		for j := 0; j < len(currenciesPerCompany); j++ {
			currrentCurrency := currenciesPerCompany[j]
			var inDenars float64
			for k := 0; k < len(valutes); k++ {
				if strings.EqualFold(currrentCurrency.Currency, valutes[k].Currency) {
					inDenars = (currrentCurrency.Amount * valutes[k].Amount)
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
