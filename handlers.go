package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pdf-golang/repository"
	"strconv"
	"strings"
)

type App struct {
	repo 		*repository.OrdersRepository
	indexTmpl 	*template.Template
	orderTmpl 	*template.Template
}

type OrdersPageData struct{
	Orders []repository.Order
	Title string
}

type OrderPageData struct{
	Order repository.Order
	Title string
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request){
	data := OrdersPageData{
		Orders: a.repo.GetAllOrders(),
		Title: "Все заказы",
	}

	if err := a.indexTmpl.Execute(w, data); err != nil{
		internalError(w, err)
	}
}

func (a *App) OrderHandler(w http.ResponseWriter, r *http.Request){
	id, err := parseID(r.URL.Path, "/order/")
	if err != nil{
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	order, ok := a.repo.GetOrderByID(id)
	if !ok{
		http.NotFound(w, r)
		return
	}

	data := OrderPageData{
		Order: order,
		Title: fmt.Sprintf("Заказ #%d", id),
	}

	if err := a.orderTmpl.Execute(w, data); err != nil{
		internalError(w, err)
	}

}

func (a *App) AllOrdersPDFHandler(w http.ResponseWriter, r *http.Request) {

	data := OrdersPageData{
		Orders: a.repo.GetAllOrders(),
		Title:  "Все заказы",
	}

	if err := writePDF(
		w,
		a.indexTmpl,
		data,
		"orders.pdf",
	); err != nil {
		internalError(w, err)
	}
}

func (a *App) OrderPDFHandler(w http.ResponseWriter, r *http.Request) {

	id, err := parseID(r.URL.Path, "/pdf/order/")
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	order, ok := a.repo.GetOrderByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	data := OrderPageData{
		Order: order,
		Title: fmt.Sprintf("Заказ #%d", id),
	}

	filename := fmt.Sprintf("order-%d.pdf", id)

	if err := writePDF(
		w,
		a.orderTmpl,
		data,
		filename,
	); err != nil {
		internalError(w, err)
	}
}


func parseID(path, prefix string) (int, error){
	idStr := strings.TrimPrefix(path, prefix)
	return strconv.Atoi(idStr)
}


func internalError(w http.ResponseWriter, err error){
	log.Println(err)

	http.Error(
		w,
		"Внутрення ошибка сервера",
		http.StatusInternalServerError,
	)
}