package crud

import (
	"errors"
	"sorotele-backend/database"

	"gorm.io/gorm"
)

// Создание нового пользователя
func CreateUser(db *gorm.DB, user database.User) {
	db.Create(&user)
}

// Получить всех пользователей
func GetAllUsers(db *gorm.DB) []database.User {
	var users []database.User
	db.Find(&users)
	return users
}

func GetAllUsersByRoleName(db *gorm.DB, role string) []database.User {
	var users []database.User
	db.Joins("JOIN roles ON roles.id = users.role_id").Where("roles.name = ?", role).Find(&users)
	return users
}

// Получить пользователя по лицевому счёту
func GetUserByAccount(db *gorm.DB, account string) (*database.User, error) {
	var user database.User
	if result := db.First(&user, "account = ?", account); result.Error != nil &&
		result.Error == gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return &user, nil
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
