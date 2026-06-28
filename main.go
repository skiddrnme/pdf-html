package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"pdf-golang/repository"
	"strconv"
)

func mul(price float64, quantity int) float64 {
	return price * float64(quantity)
}

func sumItems(items []repository.Item) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func main() {
	funcMap := template.FuncMap{
		"mul":           mul,
		"sumOrderItems": sumItems,
	}

	repo := &repository.OrdersRepository{}

	// Парсим шаблоны
	indexTmpl := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("./templates/index.html"))
	orderTmpl := template.Must(template.New("order.html").Funcs(funcMap).ParseFiles("./templates/order.html"))

	// Главная страница - список всех заказов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		orders := repo.GetAllOrders()

		data := struct {
			Orders []repository.Order
			Title  string
		}{
			Orders: orders,
			Title:  "Все заказы",
		}

		if err := indexTmpl.Execute(w, data); err != nil {
			log.Printf("Ошибка рендеринга: %v", err)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		}
	})

	// Страница конкретного заказа
	http.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/order/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Неверный ID заказа", http.StatusBadRequest)
			return
		}

		order, found := repo.GetOrderByID(id)
		if !found {
			http.Error(w, "Заказ не найден", http.StatusNotFound)
			return
		}

		data := struct {
			Order repository.Order
			Title string
		}{
			Order: order,
			Title: "Заказ #" + strconv.Itoa(id),
		}

		if err := orderTmpl.Execute(w, data); err != nil {
			log.Printf("Ошибка рендеринга: %v", err)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		}
	})

	// Эндпоинт для PDF всех заказов
	http.HandleFunc("/pdf/all", func(w http.ResponseWriter, r *http.Request) {
		orders := repo.GetAllOrders()
		data := struct {
			Orders []repository.Order
			Title  string
		}{
			Orders: orders,
			Title:  "Все заказы",
		}

		var htmlBuf bytes.Buffer
		if err := indexTmpl.Execute(&htmlBuf, data); err != nil {
			log.Printf("Ошибка рендеринга HTML: %v", err)
			http.Error(w, "Ошибка генерации PDF", http.StatusInternalServerError)
			return
		}

		pdf, err := generatePDFWithWkhtmltopdf(htmlBuf.String())
		if err != nil {
			log.Printf("Ошибка генерации PDF: %v", err)
			http.Error(w, "Ошибка генерации PDF", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=orders.pdf")
		w.Write(pdf)
	})

	

	// Статика
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Сервер запущен на :8080")
	log.Println("Доступные URL:")
	log.Println("  - http://localhost:8080/ - список заказов")
	log.Println("  - http://localhost:8080/order/102 - детали заказа")
	log.Println("  - http://localhost:8080/pdf/all - скачать PDF всех заказов")
	log.Println("  - http://localhost:8080/pdf/order/102 - скачать PDF заказа")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generatePDFWithWkhtmltopdf(htmlContent string) ([]byte, error) {
	// Временный HTML
	tmpHTML, err := os.CreateTemp("", "order-*.html")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpHTML.Name())

	if _, err := tmpHTML.WriteString(htmlContent); err != nil {
		return nil, err
	}
	tmpHTML.Close()

	// Временный PDF
	tmpPDF, err := os.CreateTemp("", "order-*.pdf")
	if err != nil {
		return nil, err
	}
	tmpPDF.Close()
	defer os.Remove(tmpPDF.Name())

	// Запуск wkhtmltopdf
	cmd := exec.Command(
		"wkhtmltopdf",
		"--enable-local-file-access", // чтобы загружались CSS и изображения
		"--encoding", "utf-8",
		tmpHTML.Name(),
		tmpPDF.Name(),
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("wkhtmltopdf: %v\n%s", err, string(output))
	}

	return os.ReadFile(tmpPDF.Name())
}