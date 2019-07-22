package main

import(
	"fmt"
	"html/template"
	"log"
	"net/http"

	"../model"
)

func main() {
	http.HandleFunc("/prices", pricesHandler)
	http.HandleFunc("/reports", reportsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func pricesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO]", r.Method, r.URL)

	err := r.ParseForm()
	if err != nil {
		return
	}

	userId := r.FormValue("uid")
	if userId == "" {
		return
	}

	t, err := template.ParseFiles("prices.html")
	if err != nil {
		return
	}

	where := fmt.Sprintf("stock_id IN (SELECT stock_id FROM favourites WHERE user_id=%s)", userId)
	prices, err := model.SelectFromPricesWhere(where)
	if err != nil {
		return
	}

	t.Execute(w, prices)
}

func reportsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO]", r.Method, r.URL)

	err := r.ParseForm()
	if err != nil {
		return
	}

	stockId := r.FormValue("sid")
	if stockId == "" {
		return
	}

	t, err := template.ParseFiles("reports.html")
	if err != nil {
		return
	}

	where := fmt.Sprintf("stock_id='%s'", stockId)
	reports, err := model.SelectFromReportsWhere(where)
	if err != nil {
		return
	}

	t.Execute(w, reports)
}
