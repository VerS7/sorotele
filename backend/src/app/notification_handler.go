package app

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"sorotele-backend/auth"
	"sorotele-backend/crud"
)

type paymentNotification struct {
	Type        string
	OperationID string
	Label       string
	Secure      string
	Currency    string
	Sender      string
	Codepro     string
	Amount      string
	Datetime    string
}

// Обработчик уведомлений от YooMoney
func (a *App) PaymentNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	msg := paymentNotification{
		Type:        r.FormValue("notification_type"),
		OperationID: r.FormValue("operation_id"),
		Amount:      r.FormValue("amount"),
		Currency:    r.FormValue("currency"),
		Datetime:    r.FormValue("datetime"),
		Sender:      r.FormValue("sender"),
		Codepro:     r.FormValue("codepro"),
		Label:       r.FormValue("label"),
	}
	secret := r.FormValue("sha1_hash")

	var unhashedSecret = []string{
		msg.Type,
		msg.OperationID,
		msg.Amount,
		msg.Currency,
		msg.Datetime,
		msg.Sender,
		msg.Codepro,
		a.PaymentController.Config.Secure,
		msg.Label,
	}
	sha1hashed := auth.HashSHA1(strings.ReplaceAll(strings.Join(unhashedSecret, "&"), " ", "%20"))

	if auth.HashToString(sha1hashed[:]) != secret {
		http.Error(w, "Failed to ensure contract", http.StatusForbidden)
		return
	}

	b, err := crud.GetUserBalanceByAccount(a.DB, msg.Label)
	if err != nil {
		log.Println("Не найден пользователь с лицевым счётом: ", msg.Label)
		return
	}
	am, err := strconv.ParseFloat(msg.Amount, 64)
	if err != nil {
		return
	}

	// Примитивная тразакция
	err = func() error {
		if err := crud.CreateHistoryAttachmentByAccount(a.DB, msg.Label, crud.HistoryData{Amount: am}); err != nil {
			return err
		}
		if err := crud.UpdateUserBalanceByAccount(a.DB, msg.Label, b.Balance+am); err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		log.Println("Не удалось обновить данные пользователя: ", msg.Label)
	}
}
