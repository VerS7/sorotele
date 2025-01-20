package crud

import (
	"log"
	"sorotele-backend/database"
	"time"

	"gorm.io/gorm"
)

type HistoryData struct {
	Amount float64 `json:"amount"`
}

type HistorySnapshot struct {
	Amount   float64   `json:"amount"`
	Datetime time.Time `json:"datetime"`
}

// Создание новой записи в истории
func CreateHistoryAttachmentByAccount(db *gorm.DB, account string, history HistoryData) error {
	user, err := GetFullUserByAccount(db, account)
	if err != nil {
		return err
	}

	if result := db.Create(&database.History{
		Amount: history.Amount,
		User:   user,
	}); result.Error != nil {
		return result.Error
	}
	return nil
}

// Запрос истории по токену
func GetHistoryByToken(db *gorm.DB, token string, limit int) ([]HistorySnapshot, error) {
	var history []database.History
	var snapshots []HistorySnapshot

	user, err := GetFullUserByToken(db, token)
	if err != nil {
		log.Println(err)
		return snapshots, err
	}

	if result := db.Where("user_id = ?", user.ID).Order("created_at desc").Limit(limit).Find(&history); result.Error != nil {
		return snapshots, result.Error
	}

	for _, elem := range history {
		snapshots = append(snapshots, HistorySnapshot{Amount: elem.Amount, Datetime: elem.CreatedAt})
	}
	return snapshots, nil
}
