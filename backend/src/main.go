package main

import (
	"errors"
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
	// Загрузка параметров env параметров
	ENV_FILE := "../../.env"
	if _, err := os.Stat(ENV_FILE); errors.Is(err, os.ErrNotExist) {
		log.Println(".env файл по пути: " + ENV_FILE + " не найден.")
	} else {
		if err := godotenv.Load(ENV_FILE); err != nil {
			log.Println(err)
		}
	}

	FRONTEND_DIR := os.Getenv("FRONTEND_PATH")

	// Инициализация приложения
	app, err := app.Init(
		app.AppConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			SSL:        os.Getenv("DB_SSL"),
		},
		email.EmailControllerConfig{
			UserEmail:    os.Getenv("EMAIL_FROM"),
			UserPassword: os.Getenv("EMAIL_PASSWORD"),
			SmtpPort:     os.Getenv("EMAIL_SMTP_PORT"),
			SmtpHost:     os.Getenv("EMAIL_SMTP_HOST"),
			Recepients:   []string{os.Getenv("EMAIL_TO")},
		},
		payment.YooMoneyConfig{
			Token:             os.Getenv("YOOMONEY_CLIENT_TOKEN"),
			Reciever:          os.Getenv("YOOMONEY_RECIEVER"),
			SuccessUrl:        os.Getenv("YOOMONEY_SUCCESS_URL"),
			Secure:            os.Getenv("YOOMONEY_ENSURE_SECRET"),
			RecievePaymentUrl: os.Getenv("YOOMONEY_PAYMENT_NOTIFICATION_URL"),
		},
	)
	if err != nil {
		log.Panicln(err)
	}

	go app.DBMigrate()

	// Инициализация админа, если таковой не существует
	go app.InitAdmin(
		auth.Credentials{
			Username: os.Getenv("ADMIN_USERNAME"),
			Password: os.Getenv("ADMIN_PASSWORD"),
		},
	)

	// Файловый сервер (Для Frontend)
	fs := http.FileServer(http.Dir(FRONTEND_DIR))

	// Эндпоинты
	//// Index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(FRONTEND_DIR, r.URL.Path)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.ServeFile(w, r, FRONTEND_DIR+"/index.html")
			return
		}
		fs.ServeHTTP(w, r)
	})
	//// Авторизация
	http.HandleFunc("/api/login", AllowCORS(app.LoginHandler))
	//// Запрос услуги на подключение
	http.HandleFunc("/api/request", AllowCORS(app.OrderHandler))
	//// Оплата
	http.HandleFunc("/api/pay", AllowCORS(app.PaymentHandler))
	//// Информация о пользователе
	http.HandleFunc("/api/user", AllowCORS(app.UserDataHandler))
	//// Доп. информация о пользователе
	http.HandleFunc("/api/user/data", AllowCORS(app.UserDynamicDataHandler))
	//// Создание нового пользователя
	http.HandleFunc("/api/user/create", AllowCORS(app.CreateUserHandler))
	//// Создание нового тарифа
	http.HandleFunc("/api/rate/create", AllowCORS(app.CreateRateHandler))
	//// Получить все тарифы
	http.HandleFunc("/api/rates", AllowCORS(app.GetRatesHandler))
	//// Уведомление от YooMoney. ВАЖНО для отображения оплаты
	http.HandleFunc("/api/payment/notification", AllowCORS(app.PaymentNotificationHandler))
	//// Списание средств с пользователя
	http.HandleFunc("/api/payment/charge", AllowCORS(app.ChargeHandler))

	// Старт сервера
	log.Println("Запуск сервера...")
	if err := http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), nil); err != nil {
		log.Panicln("Ошибка при запуске сервера:", err)
	}
}
