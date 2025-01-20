package crud

import (
	"sorotele-backend/database"

	"gorm.io/gorm"
)

// Создание новой роли
func CreateRole(db *gorm.DB, role database.Role) {
	db.Create(&role)
}

// Получение роли по названию
func GetRoleByName(db *gorm.DB, name string) (database.Role, error) {
	var role database.Role
	if result := db.Where("name = ?", name).First(&role); result.Error != nil {
		return database.Role{}, result.Error
	}
	return role, nil
}
