package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Book struct {
	Rma_id       int
	Order_id     string
	Customer_id  int
	Bin_id       string
	Owner        string
	Asin         string
	Lpn          string
	Warehouse_id string
}

var tpl *template.Template
var x string
var db *sql.DB

func init() {

	tpl = template.Must(template.ParseGlob("templates/*"))
	var err error
	db, err = sql.Open("postgres", "postgres://username:password@asist.cmtikotlnagq.us-east-1.redshift.amazonaws.com:8192/asist?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/aboutus", aboutus)
	http.HandleFunc("/sims", sims)
	http.HandleFunc("/header", header)
	http.HandleFunc("/footer", footer)
	http.HandleFunc("/test", test)
	http.HandleFunc("/results", results)
	http.HandleFunc("/meettheteam", meettheteam)
	http.HandleFunc("/excellence", excellence)
	http.HandleFunc("/query", query)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":80", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func aboutus(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "aboutus.gohtml", nil)
}

func sims(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "sims.gohtml", nil)
}

func header(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "header.gohtml", nil)
}

func footer(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "footer.gohtml", nil)
}

func results(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "results.gohtml", nil)
}

func portfolio(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "portfolio.gohtml", nil)
}

func meettheteam(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "meettheteam.gohtml", nil)
}

func excellence(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "excellence.gohtml", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT rma_id, order_id, customer_id, bin_id, owner, asin, lpn, warehouse_id FROM crt.returns_in_inv WHERE warehouse_id = 'PHX7';")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Rma_id, &bk.Order_id, &bk.Customer_id, &bk.Bin_id, &bk.Owner, &bk.Asin, &bk.Lpn, &bk.Warehouse_id) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl.ExecuteTemplate(w, "test.gohtml", bks)
}

func query(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT rma_id, order_id, customer_id, bin_id, owner, asin, lpn, warehouse_id FROM crt.returns_in_inv warehouse_id = 'PHX7' limit 100;")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Rma_id, &bk.Order_id, &bk.Customer_id, &bk.Bin_id, &bk.Owner, &bk.Asin, &bk.Lpn, &bk.Warehouse_id) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tpl.ExecuteTemplate(w, "query.gohtml", bks)

}
