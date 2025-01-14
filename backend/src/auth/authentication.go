package auth

import (
	"errors"
	"sorotele-backend/crud"

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
