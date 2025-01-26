package crud

import (
	"gorm.io/gorm"

	"sorotele-backend/database"
)

type RateData struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Создание нового тарифа
func CreateRate(db *gorm.DB, rate database.Rate) error {
	result := db.Create(&rate)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Получить все тарифы
func GetAllRates(db *gorm.DB) []RateData {
	var rates []database.Rate
	var ratesData []RateData

	db.Find(&rates)

	for _, r := range rates {
		ratesData = append(ratesData, RateData{ID: r.ID, Name: r.Name, Price: r.Price})
	}
	return ratesData
}
