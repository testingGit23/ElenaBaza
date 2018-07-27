package readingCSV

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ReadCsv(name string) [][]string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("can't open %s file", name, err.Error())
	}
	defer file.Close()

	r := csv.NewReader(file)
	//r.Comma = ','
	rowsV, err := r.ReadAll()
	if err != nil {
		log.Fatalf("can't read CSV data", err.Error())
	}
	return rowsV
}

func CurrenciesReader(name string) []Currencies {
	rows := ReadCsv(name)
	ret := make([]Currencies, len(rows)-1)
	for i, elem := range rows {
		if i == 0 {
			continue
		}
		val, err := strconv.ParseFloat(elem[1], 64)
		if err != nil {
			log.Fatalf("can't parse", err.Error())
		}
		ret[i-1] = Currencies{
			Currency: elem[0],
			Amount:   val,
		}
	}
	return ret
}

type Currencies struct {
	Currency string
	Amount   float64
}

type CurrenciesPerCompany struct {
	CompanyName              string
	CurrenciesForThisCompany []Currencies
}

func CurrenciesPerCompanyReader(name string) []CurrenciesPerCompany {
	rows := ReadCsv(name)
	ret := make([]CurrenciesPerCompany, len(rows)-1)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		var name string
		var values []Currencies

		var v Currencies

		for j, elem := range row {
			//fmt.Println(elem)
			if j == 0 {
				name = elem
				continue
			} else if j%2 == 0 {
				val, err := strconv.ParseFloat(elem, 64)
				if err != nil {
					log.Fatalf("can't parse", err.Error())
				}
				v.Amount = val
				values = append(values, v)
			} else {
				v.Currency = elem
			}

		}
		ret[i-1] = CurrenciesPerCompany{
			CompanyName:              name,
			CurrenciesForThisCompany: values,
		}

	}

	return ret
}
