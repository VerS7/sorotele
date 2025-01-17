package app

import (
	"encoding/json"
	"net/http"
	"sorotele-backend/crud"
	"sorotele-backend/payment"
	"strconv"
)

type PaymentRequest struct {
	Account  string `json:"account"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Sum      string `json:"sum"`
}

// Обработчик оплаты YouMoney
func (a *App) PaymentHandler(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
}
