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

func DrinkConverter(name string) []Drink {
	matrica := ReadCsv(name)
	ret := make([]Drink, len(matrica)-1)
	for i, elem := range matrica {
		if i == 0 {
			continue
		}
		val, err := strconv.ParseFloat(elem[1], 64)
		if err != nil {
			log.Fatalf("can't parse", err.Error())
		}
		ret[i-1] = Drink{
			Name:  elem[0],
			Price: val,
		}
	}
	return ret
}

type Drink struct {
	Name  string
	Price float64
}
