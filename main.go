package main

import (
	"api.go/pkg"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("starting server")
	r := gin.Default()
    r.Use(pkg.CorsMiddleware())
	r.POST("/calculate-payment", pkg.GetPaymentAmountPerPaymentSchedule)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
