package database

import (
	"time"

	"gorm.io/gorm"
)

// Тариф
type Rate struct {
	gorm.Model
	Name  string
	Price float64
}

// Арендованные тарифы
type RentedRate struct {
	gorm.Model
	UserID   uint
	User     User `gorm:"foreignKey:UserID"`
	Adress   string
	LastPaid time.Time
}

// Пользователь
type User struct {
	gorm.Model
	Account string
	Token   string
	Name    string
	Surname string
	RoleID  uint
	Balance float64
	Role    Role `gorm:"foreignKey:RoleID"`
}

// Тип пользователя (admin, user...)
type Role struct {
	gorm.Model
	Name string
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Rate{}, &User{}, &Role{}, &RentedRate{})
}
