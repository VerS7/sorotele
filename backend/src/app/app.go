package app

import (
	"gorm.io/gorm"

	"sorotele-backend/auth"
	"sorotele-backend/crud"
	"sorotele-backend/database"
)

type App struct {
	DB        *gorm.DB
	HashIters int
}

type AppConfig struct {
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
	SSL        string
}

// Инициализация приложения
func Init(config AppConfig) (*App, error) {
	db, err := database.DBConnect(
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.SSL,
	)

	return &App{DB: db}, err
}

func (a *App) DBMigrate() {
	database.Migrate(a.DB)
}

// Инициализация администратора. Проводится единожды
func (a *App) InitAdmin(admin auth.Credentials) {
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
