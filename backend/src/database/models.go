package database

import (
	"gorm.io/gorm"
)

// Платежи
type History struct {
	gorm.Model
	Amount float64
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}

// Тариф
type Rate struct {
	gorm.Model
	Name  string
	Price float64
}

// Арендованные тарифы
type RentedRate struct {
	gorm.Model
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
	Adress string
}

// Пользователь
type User struct {
	gorm.Model
	Account string
	Token   string
	Name    string
	Surname string
	Balance float64 `gorm:"default:0"`
	RoleID  uint
	Role    Role `gorm:"foreignKey:RoleID"`
}

// Тип пользователя (admin, user...)
type Role struct {
	gorm.Model
	Name string
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Rate{}, &User{}, &Role{}, &RentedRate{}, &History{})
}
