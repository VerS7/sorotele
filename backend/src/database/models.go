package database

import (
	"gorm.io/gorm"
)

// Платежи
type History struct {
	gorm.Model
	Amount      float64
	OperationID uint
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
}

// Тариф
type Rate struct {
	gorm.Model
	Name  string
	Price float64
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
	RateID  uint
	Rate    Rate `gorm:"foreignKey:RateID"`
}

// Тип пользователя (admin, user...)
type Role struct {
	gorm.Model
	Name string
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Rate{}, &User{}, &Role{}, &History{})
}
