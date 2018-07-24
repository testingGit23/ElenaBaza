package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)



func readValutes(name string) [][]string {
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
func checkValute(valute string, rowsV [][]string) (string, error) {

	for i, _ := range rowsV {
		if i == 0 {
			continue
		}
		if strings.EqualFold(rowsV[i][0], valute) {
			return rowsV[i][1], nil
		}

	}
	return "", errors.New("value not found")
}
func calculate(rows [][]string, rowsV [][]string) [][]string {

	for i := range rows {
		if i == 0 { //row 1 where the text is
			rows[0] = append(rows[0], "Total")

		} else {

			valute := rows[i][1]
			value, err := checkValute(valute, rowsV)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(" value =", value)
			floatValutes, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Fatal("convert data is not float1", err.Error())
			}
			floatValue, err := strconv.ParseFloat(rows[i][2], 64)
			if err != nil {
				log.Fatal("convert data is not float2", err.Error())
			}
			f := floatValue / floatValutes
			rows[i] = append(rows[i], strconv.FormatFloat(f, 'f', 5, 64))
		}
	}

	return rows
}
func write(name string, rows [][]string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("can not open %s:%s", name, err.Error())
	}
	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}
func main() {
	rowsV := readValutes("Valutes.csv")
	rows := readValutes("Companies.csv")
	calculate(rows, rowsV)
	write("new.csv", rows)
	fmt.Println("rowsV", rowsV)

}
