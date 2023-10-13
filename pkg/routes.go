package pkg

import (
	"errors"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetPaymentAmoutBody struct {
	// json tag to de-serialize json body
	PropertyPrice      float64 `json:"propertyPrice"`
	DownPayment        float64 `json:"downPayment"`
	AnnualInterestRate float64 `json:"annualInterestRate"`
	AmortizationPeriod int     `json:"amortizationPeriod"`
	PaymentSchedule    string  `json:"paymentSchedule"`
}

func GetPaymentAmountPerPaymentSchedule(c *gin.Context) {
	body := GetPaymentAmoutBody{}
	log.Println("GET /calculate-payment has been called")
	if err := c.BindJSON(&body); err != nil {
		log.Println("JSON was missing some required values")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var PaymentFreqAnum int
	switch body.PaymentSchedule {

	case "accelerated bi-weekly":
		PaymentFreqAnum = 26
	case "bi-weekly":
		PaymentFreqAnum = 24
	case "monthly":
		PaymentFreqAnum = 12
	default:
		log.Println("Non-valid payment schedule was provided")
		c.AbortWithError(http.StatusBadRequest, errors.New("Non-valid payment schedule"))
		return
	}

	principle := body.PropertyPrice - body.DownPayment
	paymentScheduleInterestRate := body.AnnualInterestRate / float64(PaymentFreqAnum)
	if body.AmortizationPeriod%5 != 0 {
		log.Println("Amortization period was not a 5 based slider")
		c.AbortWithError(http.StatusBadRequest, errors.New("Amortization period not on a 5 years increment"))
		return
	}

	var totalPayments float64 = float64(body.AmortizationPeriod * PaymentFreqAnum)
	PaymentAmountPerPaymentSchedule := principle * ((paymentScheduleInterestRate * math.Pow(1+paymentScheduleInterestRate, totalPayments)) / (math.Pow(1+paymentScheduleInterestRate, totalPayments) - 1))
	if PaymentAmountPerPaymentSchedule < 0 {
		log.Println("Payment was negative")
		c.AbortWithError(http.StatusBadRequest, errors.New("Payment was negative"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"payment_per_payment_schedule": PaymentAmountPerPaymentSchedule,
	})
}
