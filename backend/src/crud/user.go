package crud

import (
	"errors"
	"sorotele-backend/database"

	"gorm.io/gorm"
)

// Создание нового пользователя
func CreateUser(db *gorm.DB, user database.User) {
	go db.Create(&user)
}

// Получить всех пользователей
func GetAllUsers(db *gorm.DB) []database.User {
	var users []database.User
	go db.Find(&users)
	return users
}

// Получить пользователя по лицевому счёту
func GetUserByAccount(db *gorm.DB, account string) database.User {
	var user database.User
	go db.First(&user, "account = ?", account)
	return user
}

// Получить пользователя по токену доступа
func GetUserByToken(db *gorm.DB, token string) (database.User, error) {
	var user database.User
	result := db.Preload("Role").First(&user, "token = ?", token)
	if result.RowsAffected == 0 {
		return user, errors.New("пользователь с данным токеном не найден")
	}
	return user, nil
}
