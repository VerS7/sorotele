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

type UserBalance struct {
	Account string  `json:"account"`
	Balance float64 `json:"balance"`
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

// Получить полную информацию пользователя по лицевому счёту
func GetFullUserByAccount(db *gorm.DB, account string) (database.User, error) {
	var user database.User
	if result := db.First(&user, "account = ?", account); result.Error != nil &&
		result.Error == gorm.ErrRecordNotFound {
		return database.User{}, result.Error
	}

	return user, nil
}

// Получить полную информацию пользователя по токену доступа
func GetFullUserByToken(db *gorm.DB, token string) (database.User, error) {
	var user database.User

	result := db.Preload("Role").First(&user, "token = ?", token)
	if result.RowsAffected == 0 {
		return database.User{}, errors.New("пользователь с данным токеном не найден")
	}

	return user, nil
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

// Получить баланс пользователя по токену доступа
func GetUserBalanceByToken(db *gorm.DB, token string) (UserBalance, error) {
	var user database.User

	result := db.Preload("Role").First(&user, "token = ?", token)
	if result.RowsAffected == 0 {
		return UserBalance{}, errors.New("пользователь с данным токеном не найден")
	}

	return UserBalance{
		Account: user.Account,
		Balance: user.Balance,
	}, nil
}

// Получить баланс пользователя по лицевому счёту
func GetUserBalanceByAccount(db *gorm.DB, account string) (UserBalance, error) {
	var user database.User

	result := db.Preload("Role").First(&user, "account = ?", account)
	if result.RowsAffected == 0 {
		return UserBalance{}, errors.New("пользователь с данным лицевым счётом не найден")
	}

	return UserBalance{
		Account: user.Account,
		Balance: user.Balance,
	}, nil
}

// Обновить баланс пользователя по лицевому счёту
func UpdateUserBalanceByAccount(db *gorm.DB, account string, balance float64) error {
	if result := db.Model(&database.User{}).Where("account = ?", account).Update("balance", balance); result.Error != nil {
		return result.Error
	}
	return nil
}
