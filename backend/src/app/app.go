package app

import (
	"log"

	"gorm.io/gorm"

	"sorotele-backend/auth"
	"sorotele-backend/crud"
	"sorotele-backend/database"
	"sorotele-backend/email"
)

type App struct {
	DB              *gorm.DB
	HashIters       int
	EmailController *email.EmailController
}

type AppConfig struct {
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
	SSL        string
}

// Инициализация приложения
func Init(ac AppConfig, ec email.EmailControllerConfig) (*App, error) {
	db, err := database.DBConnect(
		ac.DBHost,
		ac.DBUsername,
		ac.DBPassword,
		ac.DBName,
		ac.SSL,
	)
	if err != nil {
		log.Panicln("Ошибка при подключении к базе данных: ", err)
		return nil, err
	}

	emailController, err := email.Init(ec)
	if err != nil {
		log.Panicln("Ошибка при подключении к SMTP: ", err)
		return nil, err
	}

	return &App{DB: db, EmailController: emailController}, nil
}

func (a *App) DBMigrate() {
	database.Migrate(a.DB)
}

// Инициализация администратора.
func (a *App) InitAdmin(admin auth.Credentials) {
	// Проверка, есть ли админы в системе
	admins := crud.GetAllUsersByRoleName(a.DB, "Admin")
	if len(admins) != 0 {
		return
	}

	tokenHash := auth.HashSHA256(admin.Password, a.HashIters)
	token := auth.HashToString(tokenHash[:])

	if _, err := crud.GetUserByToken(a.DB, token); err != nil {
		crud.CreateUser(a.DB, database.User{
			Account: "0",
			Token:   token,
			Name:    "Admin",
			Surname: "",
			Role:    database.Role{Name: "Admin"},
		})
	}
}
