package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	//host     = "localhost"
	//port     = "3306"
	user     = "root"
	password = "12345"
	dbName   = "demodb"
)

var tmpl = template.Must(template.ParseGlob("../form/*"))

type Payment struct {
	ID       int
	Merchant string
	Currency string
	Amount   float64
	Date     string
}

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

func home(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	rows, err := db.Query("SELECT * FROM payments")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var allPayments []Payment
	for rows.Next() {
		var p Payment
		err = rows.Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
		if err != nil {
			panic(err)
		}
		allPayments = append(allPayments, p)
	}
	//fmt.Fprintln(w, allPayments)
	tmpl.ExecuteTemplate(w, "Home", allPayments)

}

func new(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func view(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	id := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM payments WHERE paymentID=(?)", id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var p Payment
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
		if err != nil {
			panic(err)
		}
	}
	tmpl.ExecuteTemplate(w, "Show", p)
}

func edit(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	id := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM payments WHERE paymentID=(?)", id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var p Payment
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Merchant, &p.Currency, &p.Amount, &p.Date)
		if err != nil {
			panic(err)
		}
	}
	tmpl.ExecuteTemplate(w, "Edit", p)
}

func insert(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	if r.Method == "POST" {
		Merchant := r.FormValue("merchant")
		Currency := r.FormValue("currency")
		Amount := r.FormValue("amount")
		Date := r.FormValue("date")
		insForm, err := db.Prepare("INSERT INTO payments(paymentID,merchantUsername, currency, amount, dateOfPayment) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(0, Merchant, Currency, Amount, Date)
		log.Println("INSERT: Merchant: " + Merchant + " | Currency: " + Currency + " | Amount: " + Amount + " | Date: " + Date)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func update(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	if r.Method == "POST" {
		Merchant := r.FormValue("merchant")
		Currency := r.FormValue("currencies")
		Amount := r.FormValue("amount")
		Date := r.FormValue("date")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE payments SET merchantUsername=(?), currency=(?), amount=(?), dateOfPayment=(?) WHERE paymentID=(?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(Merchant, Currency, Amount, Date, id)
		log.Println("UPDATE: Merchant: " + Merchant + " | Currency: " + Currency + " | Amount: " + Amount + " | Date: " + Date)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := selectDatabase()
	id := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM payments WHERE paymentID=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/new", new)
	http.HandleFunc("/view", view)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}
