package main

import (
	"ElenaBaza/glsSolution/pkg/calculate"
	"ElenaBaza/glsSolution/pkg/readingCSV"
	"ElenaBaza/glsSolution/pkg/writingTotals"
	"fmt"
	"strconv"
)

func main() {
	writingTotals.WriteCSV("asd/Companies.csv", calculate.CalculatingPerCompani())
	matrix := readingCSV.ReadValutes("asd/Companies.csv")
	total := 0.0
	for i, row := range matrix {
		if i == 0 {
			continue
		} else {
			val, _ := strconv.ParseFloat(row[len(row)-1], 64)
			total += val
		}
	}
	fmt.Println("TOTAL:   ", total)
}
