package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

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

func Increment(c echo.Context) error {
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

func Decrement(c echo.Context) error {
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
