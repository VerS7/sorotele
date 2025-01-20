package auth

import (
	"errors"
	"sorotele-backend/crud"
	"sorotele-backend/database"

	"gorm.io/gorm"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Access struct {
	Username    string `json:"username"`
	AccessToken string `json:"token"`
}

// Проверка токена
func EnsureTokenAuth(db *gorm.DB, token string) (Access, error) {
	user, err := crud.GetUserByToken(db, token)
	if err != nil {
		return Access{}, errors.New("пользователь не найден")
	}
	return Access{Username: user.Account, AccessToken: token}, nil
}

// Проверка на администратора
func EnsureAdmin(db *gorm.DB, token string) (Access, error) {
	var founded *database.User
	for _, user := range crud.GetAllUsersByRoleName(db, "Admin") {
		if user.Token == token {
			founded = &user
			break
		}
	}

	if founded != nil {
		return Access{Username: founded.Name, AccessToken: token}, nil
	}
	return Access{}, errors.New("доступ запрещён")
}
