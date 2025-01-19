package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"sorotele-backend/app"
	"sorotele-backend/auth"
	"sorotele-backend/email"
	"sorotele-backend/payment"
)

// Обертка для добавления заголовков CORS
func AllowCORS(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,  Authentication-Token")
		if r.Method == http.MethodOptions {
			return
		}
		h(w, r)
	}
}

func main() {
	// Загрузка параметров из .env файла
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}
	// Инициализация приложения
	app, err := app.Init(
		app.AppConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			SSL:        os.Getenv("SSL"),
		},
		email.EmailControllerConfig{
			UserEmail:    os.Getenv("EMAIL_FROM"),
			UserPassword: os.Getenv("EMAIL_PASSWORD"),
			Recepients:   []string{os.Getenv("EMAIL_TO")},
			SmtpPort:     os.Getenv("EMAIL_SMTP_PORT"),
			SmtpHost:     os.Getenv("EMAIL_SMTP_HOST"),
		},
		payment.YooMoneyConfig{
			Token:      os.Getenv("YOOMONEY_CLIENT_TOKEN"),
			Reciever:   os.Getenv("YOOMONEY_RECIEVER"),
			SuccessUrl: os.Getenv("SUCCESS_URL"),
		},
	)
	if err != nil {
		panic(err)
	}

	// Инициализация админа, если таковой не существует
	go app.InitAdmin(
		auth.Credentials{
			Username: os.Getenv("ADMIN_USERNAME"),
			Password: os.Getenv("ADMIN_PASSWORD"),
		},
	)

	go app.DBMigrate()

	// Файловый сервер (Для Frontend)
	fs := http.FileServer(http.Dir("../../frontend/dist"))

	// Эндпоинты
	//// Index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join("../../frontend/dist", r.URL.Path)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.ServeFile(w, r, "../../frontend/dist/index.html")
			return
		}
		fs.ServeHTTP(w, r)
	})
	//// Авторизация
	http.HandleFunc("/api/login", AllowCORS(app.LoginHandler))
	//// Запрос услуги на подключение
	http.HandleFunc("/api/request", app.OrderHandler)
	//// Оплата
	http.HandleFunc("/api/pay", app.PaymentHandler)
	//// Информация о пользователе
	http.HandleFunc("/api/user", AllowCORS(app.UserDataHandler))
	//// Информация о тарифах пользователя
	http.HandleFunc("/api/user/rates", func(w http.ResponseWriter, r *http.Request) {})
	//// Создание нового пользователя
	http.HandleFunc("/api/user/create", func(w http.ResponseWriter, r *http.Request) {})
	//// Обновление данных пользователя
	http.HandleFunc("/api/user/update", func(w http.ResponseWriter, r *http.Request) {})
	//// Удаление пользователя
	http.HandleFunc("/api/user/delete", func(w http.ResponseWriter, r *http.Request) {})

	//// Старт сервера
	if err := http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), nil); err != nil {
		log.Panicln("Ошибка при запуске сервера:", err)
	}
}
