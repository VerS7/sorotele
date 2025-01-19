package app

import (
	"encoding/json"
	"net/http"
	"sorotele-backend/auth"
	"sorotele-backend/crud"
)

// Обработчик авторизации
func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials

	w.Header().Set("Content-Type", "application/json")

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

func (a *App) UserDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
