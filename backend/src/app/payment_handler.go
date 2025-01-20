package app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"sorotele-backend/auth"
	"sorotele-backend/crud"
	"sorotele-backend/payment"
)

type PaymentRequest struct {
	Account string `json:"account"`
	Email   string `json:"email"`
	Sum     string `json:"sum"`
}

type PaymentResponse struct {
	Service string `json:"service"`
	Link    string `json:"link"`
}

type ChargeRequest struct {
	Account string  `json:"account"`
	Amount  float64 `json:"amount"`
}

type ChargeResponse struct {
	Status   string    `json:"status"`
	Account  string    `json:"account"`
	Balance  float64   `json:"balance"`
	Amount   float64   `json:"amount"`
	Datetime time.Time `json:"datetime"`
}

// Обработчик оплаты YouMoney
func (a *App) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := crud.GetUserByAccount(a.DB, data.Account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	sum, _ := strconv.Atoi(data.Sum)

	url, err := a.PaymentController.Quickpay(payment.QuickpayForm{
		Form:        "shop",
		Target:      "Оплата SoroTele",
		Label:       user.Account,
		PaymentType: "AC",
		Sum:         sum,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&PaymentResponse{Service: "Юmoney", Link: url.String()})
}

// Обработчик списания денег со счета
func (a *App) ChargeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authentication-Token")

	// Проверка токена доступа
	if len(token) == 0 {
		http.Error(w, "Failed to access", http.StatusForbidden)
		return
	}
	if _, err := auth.EnsureAdmin(a.DB, token); err != nil {
		http.Error(w, "Failed to access", http.StatusForbidden)
		return
	}

	var data ChargeRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usrb, err := crud.GetUserBalanceByAccount(a.DB, data.Account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Если баланс пользователя меньше нуля -> отказ
	if usrb.Balance < 0 {
		json.NewEncoder(w).Encode(&ChargeResponse{
			Status:   "denied",
			Account:  data.Account,
			Balance:  usrb.Balance,
			Amount:   data.Amount,
			Datetime: time.Now(),
		})
		return
	}

	charge := 0 - data.Amount

	if err := crud.CreateHistoryAttachmentByAccount(
		a.DB,
		data.Account,
		crud.HistoryData{Amount: charge},
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	crud.UpdateUserBalanceByAccount(a.DB, data.Account, usrb.Balance+charge)

	json.NewEncoder(w).Encode(&ChargeResponse{
		Status:   "success",
		Account:  data.Account,
		Balance:  usrb.Balance + charge,
		Amount:   data.Amount,
		Datetime: time.Now(),
	})
}
