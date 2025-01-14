package database

import (
	"time"

	"gorm.io/gorm"
)

// Тариф
type Rate struct {
	gorm.Model
	Name  string
	Price float32
}

// Арендованные тарифы
type RentedRate struct {
	gorm.Model
	UserID   uint
	User     User `gorm:"foreignKey:UserID"`
	Adress   string
	Balance  float32
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
