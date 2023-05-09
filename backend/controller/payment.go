package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentController struct{}

func NewPaymentController() *PaymentController {
	return &PaymentController{}
}

func (p *PaymentController) CreatePayment(c echo.Context) error {

	// just a simulation
	response := map[string]string{"message": "Płatność została zrealizowana"}

	return c.JSON(http.StatusOK, response)
}
