package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"html/template"
)



func selectDatabase() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "darko123"
	dbName := "goblog"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db

}
type Currencies struct {
	Currency string
	InDenars float64
}
var tmpl = template.Must(template.ParseGlob("form/*"))

func Valutes(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	rows, err := db.Query("SELECT * FROM currencies")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var allValutes []Currencies
	for rows.Next() {
		var c Currencies
		err = rows.Scan(&c.Currency, &c.InDenars)
		if err != nil {
			panic(err)
		}
		allValutes = append(allValutes, c)
	}
	// fmt.Fprintln(w, allValutes)

	tmpl.ExecuteTemplate(w, "Valutes", allValutes)

}

func selectAllCurrencies(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT currency FROM currencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []string
	for rows.Next() {
		var temp string
		err = rows.Scan(&temp)
		if err != nil {
			return nil, err
		}
		ret = append(ret, temp)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func selectAllMerchants(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT merchantUsername FROM merchants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []string
	for rows.Next() {
		var temp string
		err = rows.Scan(&temp)
		if err != nil {
			return nil, err
		}
		ret = append(ret, temp)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func totals(valutes []string, merchants []string, db *sql.DB) (map[string]float64, error) {
	ret := make(map[string]float64)
	for _, merchant := range merchants {
		sumPerMerchant := 0.0
		for _, inDenars := range valutes {
			var sumFromQuerry float64
			sumPerCurrencyForMerchant, err := db.Query("SELECT SUM(amount) FROM payments WHERE merchantUsername=(?) AND currency=(?)", merchant, inDenars)
			if err != nil {
				return nil, err
			}

			for sumPerCurrencyForMerchant.Next() {
				var temp float64
				err = sumPerCurrencyForMerchant.Scan(&temp)
				if err != nil {
					temp = 0
				}
				sumFromQuerry = temp
			}
			err = sumPerCurrencyForMerchant.Err()
			if err != nil {
				return nil, err
			}
			defer sumPerCurrencyForMerchant.Close()
			var InDenars float64
			currencyInDenars, err := db.Query("SELECT inDenars FROM currencies WHERE currency=(?)", inDenars)
			if err != nil {
				return nil, err
			}

			for currencyInDenars.Next() {
				var temp float64
				err = currencyInDenars.Scan(&temp)
				if err != nil {
					return nil, err
				}
				InDenars = temp
			}
			err = currencyInDenars.Err()
			if err != nil {
				return nil, err
			}
			defer currencyInDenars.Close()
			pom := sumFromQuerry * InDenars
			sumPerMerchant = sumPerMerchant + pom
		}
		ret[merchant] = sumPerMerchant
	}
	return ret, nil
}



func main() {
//	dataBase := selectDatabase()
	//currencies, _ := selectAllCurrencies(dataBase)

	//merchantsUsernames, _ := selectAllMerchants(dataBase)

	//totalsMap, _ := totals(currencies, merchantsUsernames, dataBase)

	//fmt.Println(currencies)
	//fmt.Println(merchantsUsernames)
	//fmt.Println(totalsMap)
	log.Println("Server started on: http://localhost:3606")
	http.HandleFunc("/", Valutes)
	http.ListenAndServe(":3606", nil)
}
