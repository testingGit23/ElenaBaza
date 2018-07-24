package readingCSV

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadValutes(name string) [][]string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("can't open %s file", name, err.Error())
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	rowsV, err := r.ReadAll()
	if err != nil {
		log.Fatalf("can't read CSV data", err.Error())
	}
	return rowsV
}

func CurrenciesConverter(matrica [][]string) []Currencies {
	ret := make([]Currencies, len(matrica)-1)
	for i, elem := range matrica {
		if i == 0 {
			continue
		}
		val, err := strconv.ParseFloat(elem[1], 64)
		if err != nil {
			log.Fatalf("can't parse", err.Error())
		}
		ret[i-1] = Currencies{
			Currency: elem[0],
			InDenars: val,
		}
	}
	return ret
}

type Currencies struct {
	Currency string
	InDenars float64
}

type CurrenciesPerCompany struct {
	CompanyName              string
	CurrenciesForThisCompany []Currencies
}

func CurrenciesPerCompanyConverter(matrica [][]string) []CurrenciesPerCompany {
	ret := make([]CurrenciesPerCompany, len(matrica)-1)
	for i, row := range matrica {
		if i == 0 {
			continue
		}
		var name string
		var values []Currencies
		var value float64
		niza := fmt.Sprintln(row)
		rows := strings.Split(niza, ";")
		for j, elem := range rows {
			fmt.Println(elem)
			if j == 0 {
				name = string(elem[j])
				continue
			}
			if j%2 != 0 {
				val, err := strconv.ParseFloat(string(elem[j]), 64)
				if err != nil {
					log.Fatalf("can't parse", err.Error())
				}
				value = val
			}
			v := Currencies{
				Currency: string(elem[j]),
				InDenars: value,
			}
			values = append(values, v)
		}
		ret[i-1] = CurrenciesPerCompany{
			CompanyName:              name,
			CurrenciesForThisCompany: values,
		}

	}

	return ret
}
