package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPaymentAmountPerPaymentSchedule_ValidRequest_Monthly(t *testing.T) {
	r := gin.Default()
	r.POST("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

	requestBody := `{
        "propertyPrice": 200000,
        "downPayment": 40000,
        "annualInterestRate": 0.04,
        "amortizationPeriod": 20,
        "paymentSchedule": "monthly"
    }`

	req, _ := http.NewRequest("POST", "/calculate-payment", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedResponse := `{"payment_per_payment_schedule":969.5685268790528}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, w.Body.String())
	}
}

func TestGetPaymentAmountPerPaymentSchedule_ValidRequest_Biweekly(t *testing.T) {
	r := gin.Default()
	r.POST("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

	requestBody := `{
        "propertyPrice": 200000,
        "downPayment": 40000,
        "annualInterestRate": 0.04,
        "amortizationPeriod": 20,
        "paymentSchedule": "bi-weekly"
    }`

	req, _ := http.NewRequest("POST", "/calculate-payment", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedResponse := `{"payment_per_payment_schedule":484.5210226285538}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, w.Body.String())
	}
}

func TestGetPaymentAmountPerPaymentSchedule_ValidRequest_Acc_Biweekly(t *testing.T) {
	r := gin.Default()
	r.POST("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

	requestBody := `{
        "propertyPrice": 200000,
        "downPayment": 40000,
        "annualInterestRate": 0.04,
        "amortizationPeriod": 20,
        "paymentSchedule": "accelerated bi-weekly"
    }`

	req, _ := http.NewRequest("POST", "/calculate-payment", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedResponse := `{"payment_per_payment_schedule":447.23147838283364}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, w.Body.String())
	}
}

func TestGetPaymentAmountPerPaymentSchedule_InvalidRequest(t *testing.T) {
	r := gin.Default()
	r.POST("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

	// Invalid request with missing fields
	requestBody := `{
        "propertyPrice": 200000,
        "downPayment": 40000
        // Missing other fields
    }`

	req, _ := http.NewRequest("POST", "/calculate-payment", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetPaymentAmountPerPaymentSchedule_InvalidPayment(t *testing.T) {
	r := gin.Default()
	r.POST("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

	requestBody := `{
        "propertyPrice": 200000,
        "downPayment": 40000,
        "annualInterestRate": 0.04,
        "amortizationPeriod": 20,
        "paymentSchedule": "NULL"
    }`

	req, _ := http.NewRequest("POST", "/calculate-payment", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetPaymentAmountPerPaymentSchedule_InvalidRequest_Negative_Payment(t *testing.T) {
	r := gin.Default()
	r.POST("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

	requestBody := `{
        "propertyPrice": -200000,
        "downPayment": 40000,
        "annualInterestRate": 0.04,
        "amortizationPeriod": 20,
        "paymentSchedule": "accelerated bi-weekly"
    }`

	req, _ := http.NewRequest("POST", "/calculate-payment", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
}
