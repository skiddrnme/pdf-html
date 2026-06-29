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
```


## 🌐 Доступные маршруты

### 🖥️ UI (Web интерфейс)

- `GET /`  
  📄 Список всех заказов

- `GET /order/{id}`  
  📄 Детальная информация по конкретному заказу

---

### 📦 PDF экспорт

- `GET /pdf/all`  
  📥 Скачать PDF со всеми заказами

- `GET /pdf/order/{id}`  
  📥 Скачать PDF конкретного заказа

---

## 🧪 Пример данных

В проекте используются mock-заказы:

- 📊 5 заказов
- 🛒 Товары:
  - ноутбуки
  - мониторы
  - смартфоны
  - планшеты
  - аксессуары

---

## 🧠 Архитектура

Проект разделён на логические слои:

- 🧩 `handlers` — HTTP-логика (presentation layer)
- 🗄 `repository` — работа с данными (data layer)
- 🧮 `utils` — вспомогательные функции (business logic helpers)
- 📄 `pdf generator` — генерация PDF через wkhtmltopdf
- 🎨 `templates` — HTML представление

---

## 🚀 Возможные улучшения

- 🐘 Подключить PostgreSQL вместо mock data
- 🐳 Добавить Docker
- 📊 Добавить логирование (zap / slog)
- ⚙️ Конфигурация через env (viper)
- 🧭 Router (chi / gorilla/mux)
- 🔐 Авторизация пользователей
- ⚡ Кеширование PDF

---

## 📝 Примечание

Проект учебный и демонстрирует:

- работу с HTTP в Go
- HTML templates
- генерацию PDF через внешние инструменты
- базовую архитектуру backend-приложения
