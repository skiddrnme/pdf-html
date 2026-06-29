# PDF Orders Service (Go)

Простое веб-приложение на Go для просмотра заказов и генерации PDF-отчетов с использованием `wkhtmltopdf`.

## 🚀 Возможности

- Просмотр списка заказов
- Просмотр деталей конкретного заказа
- Генерация PDF всех заказов
- Генерация PDF одного заказа
- Шаблоны на HTML + Go templates
- Расчеты сумм через template functions
- Статические файлы (CSS/JS)

---

## 🧱 Технологии

- Go (net/http)
- html/template
- wkhtmltopdf
- MVC-подобная структура (упрощённая)
- in-memory repository (mock data)

---

## 📁 Структура проекта

pdf-golang/
├── main.go
├── handlers.go
├── pdf.go
├── utils/
├── repository/
├── templates/
└── static/

## ▶️ Запуск проекта

### 1. Установи зависимости

Убедись, что установлен Go:

```bash
go version

2. Установи wkhtmltopdf
Ubuntu / WSL:
sudo apt install wkhtmltopdf

macOS:
brew install wkhtmltopdf

Windows:
Скачать:
https://wkhtmltopdf.org/downloads.html

# 3. Запуск приложения
go run .
или
go build .
./pdf-golang


## Доступные маршруты
- UI
- GET / — список всех заказов
- GET /order/{id} — детали заказа
- PDF
- GET /pdf/all — скачать PDF всех заказов
- GET /pdf/order/{id} — скачать PDF конкретного заказа


🧪 Пример данных
В проекте используются mock-заказы:
5 заказов
товары: ноутбуки, мониторы, смартфоны и т.д.


##Архитектура
## Проект разделён на слои:

- handlers — HTTP-логика
- repository — работа с данными
- utils — вспомогательные функции
- pdf generator — генерация PDF через wkhtmltopdf
- templates — HTML представление


##Возможные улучшения
- Подключить PostgreSQL вместо mock data
- Добавить Docker
- Добавить логирование (zap/logrus)
- Добавить конфигурацию (env/viper)
- Разделить роутинг (gorilla/mux или chi)
- Добавить авторизацию
- Кеширование PDF


## Примечание
## Проект учебный, демонстрирует:

- работу с HTTP в Go
- шаблоны HTML
- генерацию PDF через внешний инструмент
- базовую архитектуру приложения
```
