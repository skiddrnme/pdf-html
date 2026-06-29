package main

import (
	"html/template"
	"log"
	"net/http"
	"pdf-golang/repository"
	"pdf-golang/utils"
)

func main() {
	funcMap := template.FuncMap{
		"mul":           utils.Mul,
		"sumOrderItems": utils.SumItems,
	}

	app := &App{
		repo: repository.NewOrdersRepository(),
		indexTmpl: template.Must(
			template.New("index.html").Funcs(funcMap).ParseFiles("./templates/index.html"),
		),
		orderTmpl: template.Must(
			template.New("order.html").Funcs(funcMap).ParseFiles("./templates/order.html"),
		),
	}

	http.HandleFunc("/", app.IndexHandler)
	http.HandleFunc("/order/", app.OrderHandler)

	http.HandleFunc("/pdf/all", app.AllOrdersPDFHandler)
	http.HandleFunc("/pdf/order/", app.OrderPDFHandler)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static")),
		),
	)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}