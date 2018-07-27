package writingTotals

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func write(name string, rows [][]string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("can not open %s:%s", name, err.Error())
	}
	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}

func WriteCSV(name string, totalsPerCompany map[string]string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("can't open %s file", name, err.Error())
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf("can't read CSV data", err.Error())
	}
	newRows, total := changingRows(rows, totalsPerCompany)
	write("asd/Companies.csv", newRows)
	fmt.Println("TOTAL:   ", total)
}

func changingRows(rows [][]string, totalsPerCompany map[string]string) ([][]string, float64) {

	lenth := len(rows[0])
	ret := rows
	if rows[0][lenth-1] == "Total" {
		for i, row := range rows {
			ret[i] = row[:lenth-1]
		}
	}
	total := 0.0
	for i := 0; i < len(ret); i++ {
		if i == 0 {
			ret[0] = append(ret[0], "Total")
		} else {
			for k, v := range totalsPerCompany {
				if strings.EqualFold(ret[i][0], k) {
					ret[i] = append(ret[i], v)
					val, _ := strconv.ParseFloat(v, 64)
					total += val
				}
			}
		}
	}

	return ret, total
}
