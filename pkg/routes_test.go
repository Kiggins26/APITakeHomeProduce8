package pkg

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/gin-gonic/gin"
)

func TestGetPaymentAmountPerPaymentSchedule_ValidRequest(t *testing.T) {
    r := gin.Default()
    r.GET("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

    requestBody := `{
        "propertyPrice": 200000,
        "downPayment": 40000,
        "annualInterestRate": 0.04,
        "amortizationPeriod": 20,
        "paymentSchedule": "monthly"
    }`

    req, _ := http.NewRequest("GET", "/calculate-payment", strings.NewReader(requestBody))
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

func TestGetPaymentAmountPerPaymentSchedule_InvalidRequest(t *testing.T) {
    r := gin.Default()
    r.GET("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

    // Invalid request with missing fields
    requestBody := `{
        "propertyPrice": 200000,
        "downPayment": 40000
        // Missing other fields
    }`

    req, _ := http.NewRequest("GET", "/calculate-payment", strings.NewReader(requestBody))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
    }
}

func TestGetPaymentAmountPerPaymentSchedule_InvalidPayment(t *testing.T) {
    r := gin.Default()
    r.GET("/calculate-payment", GetPaymentAmountPerPaymentSchedule)

    requestBody := `{
        "propertyPrice": 200000,
        "downPayment": 40000,
        "annualInterestRate": 0.04,
        "amortizationPeriod": 20,
        "paymentSchedule": "NULL"
    }`

    req, _ := http.NewRequest("GET", "/calculate-payment", strings.NewReader(requestBody))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
    }
}
