package app

import (
	"encoding/json"
	"net/http"
	"sorotele-backend/email"
)

type Order struct {
	FullName string `json:"fullName"`
	Contacts string `json:"contacts"`
	Message  string `json:"message"`
}

// Обработчик отправки email-ов с заказами
func (a *App) OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data Order
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body :=
		"Поступил новый запрос на обработку\n" +
			"ФИО: " +
			data.FullName + "\n" +
			"Контакты: " +
			data.Contacts + "\n" +
			"Дополнительно: " +
			data.Message

	a.EmailController.SendMessage(
		email.Message{
			To:      a.EmailController.Config.Recepients,
			Subject: "Новый запрос на обработку",
			Body:    body,
		})
}
