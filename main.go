package main

import (
//    "log"
    "api.go/pkg"

    "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/calculate-payment", pkg.GetPaymentAmountPerPaymentSchedule)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
