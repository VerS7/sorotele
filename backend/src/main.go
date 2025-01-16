package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"sorotele-backend/app"
	"sorotele-backend/auth"
	"sorotele-backend/email"
)

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

	app.EmailController.SendMessage(
		email.Message{
			To:      []string{"thegoversus@gmail.com"},
			Subject: "Test 1",
			Body:    "Ещё один тест",
		},
	)

	fs := http.FileServer(http.Dir("../../frontend/dist"))

	// Эндпоинты
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join("../../frontend/dist", r.URL.Path)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.ServeFile(w, r, "../../frontend/dist/index.html")
			return
		}
		fs.ServeHTTP(w, r)
	})
	http.HandleFunc("/api/login", app.LoginHandler)
	http.HandleFunc("/api/request", app.OrderHandler)
	http.HandleFunc("/api/pay", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/api/user/create", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/api/user/update", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/api/user/delete", func(w http.ResponseWriter, r *http.Request) {})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
