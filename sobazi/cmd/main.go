package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	//host     = "localhost"
	//port     = "3306"
	user     = "root"
	password = "12345"
	dbName   = "demodb"
)

func selectDatabase() (db *sql.DB) {
	db, err := sql.Open("mysql",
		user+":"+password+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(db)
	//defer db.Close()
	return db
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
	dataBase := selectDatabase()
	currencies, _ := selectAllCurrencies(dataBase)

	merchantsUsernames, _ := selectAllMerchants(dataBase)

	totalsMap, _ := totals(currencies, merchantsUsernames, dataBase)

	fmt.Println(currencies)
	fmt.Println(merchantsUsernames)
	fmt.Println(totalsMap)
	//fmt.Println(53624.3 * 53.65)
}
