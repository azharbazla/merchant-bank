package controllers

import (
	"encoding/json"
	"merchant-bank/services"
	"merchant-bank/utils"
	"net/http"
	"strings"
)

func GetAllHistory(w http.ResponseWriter) {
	history, err := services.GetHistory()
	if err != nil {
		http.Error(w, "Error fetching history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(history)
	if err != nil {
		return
	}
}

func GetCustomerHistory(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	customerID, err := utils.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	history, err := services.GetCustomerHistory(customerID)
	if err != nil {
		http.Error(w, "Error fetching customer history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(history)
	if err != nil {
		return
	}
}
