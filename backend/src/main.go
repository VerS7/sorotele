package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"sorotele-backend/app"
	"sorotele-backend/auth"
)

func main() {
	// Загрузка параметров из .env файла
	if err := godotenv.Load("../.env"); err != nil {
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
	)
	if err != nil {
		panic(err)
	}

	go app.InitAdmin(
		auth.Credentials{
			Username: os.Getenv("ADMIN_USERNAME"),
			Password: os.Getenv("ADMIN_PASSWORD"),
		},
	)

	go app.DBMigrate()

	// Эндпоинты
	http.HandleFunc("/api/login", app.LoginHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

// database.Migrate(db)

// db.Create(&database.Rate{Name: "test", Price: 322})
// var rate database.Rate
// db.Take(&rate)
// fmt.Println(rate)
// db.Take(&rate)
// fmt.Println(rate)
// db.Take(&rate)
// fmt.Println(rate)
// db.Take(&rate)
// fmt.Println(rate)

// if err != nil {
// 	panic(err)
// }
