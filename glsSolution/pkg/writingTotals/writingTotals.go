package writingTotals

import (
	"encoding/csv"
	"log"
	"os"
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
	//r.Comma = ','
	rowsV, err := r.ReadAll()
	red := rowsV
	if err != nil {
		log.Fatalf("can't read CSV data", err.Error())
	}
	new := changingRow(red, totalsPerCompany)
	write("New.csv", new)
}

func changingRow(matrica [][]string, totalsPerCompany map[string]string) [][]string {
	var ret [][]string
	ret = matrica
	for i := 0; i < len(ret); i++ {
		if i == 0 {
			ret[0] = append(ret[0], "Total")
		} else {
			for k, v := range totalsPerCompany {
				if strings.EqualFold(ret[i][0], k) {
					ret[i] = append(ret[i], v)
				}
			}
		}
	}
	return ret
}
