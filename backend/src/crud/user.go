package crud

import (
	"errors"
	"sorotele-backend/database"

	"gorm.io/gorm"
)

type UserData struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Role    string `json:"role"`
}

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
func GetUserByAccount(db *gorm.DB, account string) (UserData, error) {
	var user database.User
	if result := db.First(&user, "account = ?", account); result.Error != nil &&
		result.Error == gorm.ErrRecordNotFound {
		return UserData{}, result.Error
	}

	return UserData{
		Account: user.Account,
		Name:    user.Name,
		Surname: user.Surname,
		Role:    user.Role.Name,
	}, nil
}

// Получить пользователя по токену доступа
func GetUserByToken(db *gorm.DB, token string) (UserData, error) {
	var user database.User

	result := db.Preload("Role").First(&user, "token = ?", token)
	if result.RowsAffected == 0 {
		return UserData{}, errors.New("пользователь с данным токеном не найден")
	}

	return UserData{
		Account: user.Account,
		Name:    user.Name,
		Surname: user.Surname,
		Role:    user.Role.Name,
	}, nil
}
