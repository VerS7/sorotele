package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnect(host string, user string, password string, name string, sslmode string) (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, name, sslmode)
	return gorm.Open(postgres.Open(conn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}
