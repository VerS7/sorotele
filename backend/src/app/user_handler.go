package app

import (
	"encoding/json"
	"net/http"

	"sorotele-backend/auth"
	"sorotele-backend/crud"
	"sorotele-backend/database"

	"gorm.io/gorm"
)

type DynamicData struct {
	Balance float64                `json:"balance"`
	History []crud.HistorySnapshot `json:"history"`
}

type CreateUserData struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Role     string `json:"role"`
	RateID   uint   `json:"rateID"`
}

type CreateUserResponse struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Token   string `json:"token"`
}

type CreateRateData struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Обработчик авторизации
func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid Credentials structure", http.StatusBadRequest)
		return
	}

	token := auth.HashSHA256(creds.Password, a.HashIters)
	access, err := auth.EnsureTokenAuth(a.DB, auth.HashToString(token[:]))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewEncoder(w).Encode(access); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Обработчик данных пользователя
func (a *App) UserDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authentication-Token")
	if len(token) == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err := auth.EnsureTokenAuth(a.DB, token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userData, err := crud.GetUserByToken(a.DB, token)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(userData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Обработчик дополнительных данных пользователя (баланс, история...)
func (a *App) UserDynamicDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authentication-Token")
	if len(token) == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err := auth.EnsureTokenAuth(a.DB, token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	usbl, err := crud.GetUserBalanceByToken(a.DB, token)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	history, err := crud.GetHistoryByToken(a.DB, token, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(DynamicData{Balance: usbl.Balance, History: history}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Создание нового пользователя
func (a *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authentication-Token")
	if len(token) == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if _, err := auth.EnsureAdmin(a.DB, token); err != nil {
		http.Error(w, "Failed to access", http.StatusForbidden)
		return
	}

	var data CreateUserData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userTokenByte := auth.HashSHA256(data.Password, a.HashIters)
	userToken := auth.HashToString(userTokenByte[:])
	account := auth.GenerateAccount("sr", 16)

	role, err := crud.GetRoleByName(a.DB, data.Role)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := crud.CreateUser(a.DB, database.User{
				Account: account,
				Name:    data.Name,
				Surname: data.Surname,
				Token:   userToken,
				RateID:  data.RateID,
				Role:    database.Role{Name: "User"},
				Balance: 0,
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := crud.CreateUser(a.DB, database.User{
			Account: account,
			Name:    data.Name,
			Surname: data.Surname,
			Token:   userToken,
			RateID:  data.RateID,
			RoleID:  role.ID,
			Balance: 0,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := json.NewEncoder(w).Encode(&CreateUserResponse{
		Account: account,
		Name:    data.Name,
		Surname: data.Surname,
		Token:   userToken,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Создание нового тарифа
func (a *App) CreateRateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authentication-Token")
	if len(token) == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if _, err := auth.EnsureAdmin(a.DB, token); err != nil {
		http.Error(w, "Failed to access", http.StatusForbidden)
		return
	}

	var data CreateRateData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := crud.CreateRate(a.DB, database.Rate{
		Name:  data.Name,
		Price: data.Price,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Получение всех тарифов
func (a *App) GetRatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authentication-Token")
	if len(token) == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if _, err := auth.EnsureAdmin(a.DB, token); err != nil {
		http.Error(w, "Failed to access", http.StatusForbidden)
		return
	}

	if err := json.NewEncoder(w).Encode(crud.GetAllRates(a.DB)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
