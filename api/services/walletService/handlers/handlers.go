package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// NewAccount creates a new account by name and phone number and set the balance to 0
func NewAccount(c echo.Context) error {
	type request struct {
		fullName    string
		phoneNumber string
	}

	req := &request{}
	bindErr := c.Bind(req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	err := NewAcc(req.fullName, req.phoneNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "unable to create account")
	}

	return c.JSON(http.StatusCreated, "account created")
}

// Balance will return the balance of the wallet by phone number
func Balance(c echo.Context) error {
	type request struct {
		phoneNumber string
	}

	req := &request{}
	bindErr := c.Bind(req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	balance, err := GetBalance(req.phoneNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone number")
	}

	return c.JSON(http.StatusOK, balance)
}

// Charge will charge the wallet by giving amount and phone number
func Charge(c echo.Context) error {
	type request struct {
		amount      int64
		phoneNumber string
	}

	req := &request{}
	bindErr := c.Bind(req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	wallet, err := GetWallet(req.phoneNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone number")
	}

	wallet.IncreaseBalance(req.amount)

	return c.JSON(http.StatusOK, "your wallet balance has been updated")
}

// Use will use the wallet for given phone number and amount
func Use(c echo.Context) error {
	type request struct {
		amount      int64
		phoneNumber string
	}

	req := &request{}
	bindErr := c.Bind(req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	wallet, err := GetWallet(req.phoneNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone number")
	}

	wallet.DecreaseBalance(req.amount)

	return c.JSON(http.StatusOK, "your wallet balance has been updated")

}
