package crud

import (
	"errors"
	"sorotele-backend/database"
	"time"

	"gorm.io/gorm"
)

type HistoryData struct {
	Amount      float64 `json:"amount"`
	OperationID uint    `json:"operation_id"`
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
		Amount:      history.Amount,
		OperationID: history.OperationID,
		User:        user,
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

// Запрос истории по айдишнику
func GetHistoryByOperationID(db *gorm.DB, operationID uint) (*database.History, error) {
	var history database.History

	if result := db.Where("operation_id = ?", operationID).Find(&history); result.Error != nil || result.RowsAffected == 0 {
		if result.Error != nil {
			return nil, result.Error
		} else {
			return nil, errors.New("history with this operation_id not found")
		}
	}
	return &history, nil
}
