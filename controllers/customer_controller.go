package controllers

import (
	"encoding/json"
	"merchant-bank/models"
	"merchant-bank/services"
	"merchant-bank/utils"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := services.LoginCustomer(customer); err != nil {
		http.Error(w, "Customer not found", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(customer.ID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	utils.LogHistory("Login successful", customer.ID)
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
}

func Payment(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	customerID, err := utils.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var paymentRequest struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := services.ProcessPayment(customerID, paymentRequest.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.LogHistory("Payment successful", customerID)
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	customerID, err := utils.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := services.LogoutCustomer(customerID); err != nil {
		http.Error(w, "Logout failed", http.StatusBadRequest)
		return
	}

	utils.LogHistory("Logout successful", customerID)
	w.WriteHeader(http.StatusOK)
}
