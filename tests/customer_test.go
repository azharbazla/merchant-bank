package tests

import (
	"bytes"
	"encoding/json"
	"merchant-bank/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	requestBody := []byte(`{"email": "azhar@gmail.com", "password": "password"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	if len(rr.Header().Get("Authorization")) == 0 {
		t.Error("Expected JWT token in Authorization header")
	}

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode JSON response")
	}

	if response["id"] == nil || response["name"] == nil {
		t.Error("Expected customer data in response")
	}
}

func TestLoginFailure(t *testing.T) {
	requestBody := []byte(`{"email": "bazla@gmail.com", "password": "ppppp"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Expected status code %v, got %v", http.StatusUnauthorized, status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode JSON response")
	}

	if response["error"] == nil {
		t.Error("Expected error message in response")
	}
}

func TestPaymentUnauthorized(t *testing.T) {
	req, err := http.NewRequest("POST", "/payment", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Payment)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Expected status code %v, got %v", http.StatusUnauthorized, status)
	}
}

func TestPaymentAuthorized(t *testing.T) {
	loginRequestBody := []byte(`{"email": "azhar@gmail.com", "password": "password"}`)
	loginReq, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginRequestBody))
	if err != nil {
		t.Fatal(err)
	}
	loginRR := httptest.NewRecorder()
	loginHandler := http.HandlerFunc(controllers.Login)
	loginHandler.ServeHTTP(loginRR, loginReq)

	jwtToken := loginRR.Header().Get("Authorization")
	if jwtToken == "" {
		t.Fatal("Expected JWT token after login")
	}

	paymentRequestBody := []byte(`{"merchant_id": "1", "amount": 100}`)
	paymentReq, err := http.NewRequest("POST", "/payment", bytes.NewBuffer(paymentRequestBody))
	if err != nil {
		t.Fatal(err)
	}
	paymentReq.Header.Set("Authorization", jwtToken)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Payment)

	handler.ServeHTTP(rr, paymentReq)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode JSON response")
	}

	if response["message"] == nil || response["message"] != "Payment successful" {
		t.Error("Expected 'Payment successful' message in response")
	}
}

func TestLogout(t *testing.T) {
	loginRequestBody := []byte(`{"email": "azhar@gmail.com", "password": "password"}`)
	loginReq, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginRequestBody))
	if err != nil {
		t.Fatal(err)
	}
	loginRR := httptest.NewRecorder()
	loginHandler := http.HandlerFunc(controllers.Login)
	loginHandler.ServeHTTP(loginRR, loginReq)

	jwtToken := loginRR.Header().Get("Authorization")
	if jwtToken == "" {
		t.Fatal("Expected JWT token after login")
	}

	logoutReq, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	logoutReq.Header.Set("Authorization", jwtToken)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Logout)

	handler.ServeHTTP(rr, logoutReq)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}

	reqAfterLogout, err := http.NewRequest("POST", "/payment", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqAfterLogout.Header.Set("Authorization", jwtToken)

	rrAfterLogout := httptest.NewRecorder()
	handler.ServeHTTP(rrAfterLogout, reqAfterLogout)

	if status := rrAfterLogout.Code; status != http.StatusUnauthorized {
		t.Errorf("Expected status code %v after logout, got %v", http.StatusUnauthorized, status)
	}
}
