package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readValutes(name string) []ValuteToDenars {
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
	ret := make([]ValuteToDenars, len(rowsV)-1)
	for i, elem := range rowsV {
		if i == 0 {
			continue
		}
		val, err := strconv.ParseFloat(elem[1], 32)
		if err != nil {
			log.Fatalf("can't parse data", err.Error())
		}
		ret[i-1] = ValuteToDenars{
			valute:   elem[0],
			inDenars: val,
		}
	}
	return ret
}

func main() {
	fmt.Println(readValutes("Valutes.csv"))
}

type ValuteToDenars struct {
	valute   string
	inDenars float64
}
