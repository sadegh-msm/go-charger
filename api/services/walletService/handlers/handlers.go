package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// NewAccount creates a new account by name and phone number and set the balance to 0
func NewAccount(c echo.Context) error {
	type request struct {
		FullName    string `json:"fullName"`
		PhoneNumber string `json:"phoneNumber"`
	}

	req := request{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	log.Println(req.FullName, req.PhoneNumber)
	log.Println("fuck")

	err := NewAcc(req.FullName, req.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "unable to create account")
	}

	return c.JSON(http.StatusCreated, "account created")
}

// Balance will return the balance of the wallet by phone number
func Balance(c echo.Context) error {
	type request struct {
		PhoneNumber string `json:"phoneNumber"`
	}

	req := request{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	balance, err := GetBalance(req.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone number")
	}

	return c.JSON(http.StatusOK, balance)
}

// Charge will charge the wallet by giving amount and phone number
func Charge(c echo.Context) error {
	type request struct {
		Amount      int64  `json:"amount"`
		PhoneNumber string `json:"phoneNumber"`
	}

	req := request{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	wallet, err := GetWallet(req.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone number")
	}

	wallet.IncreaseBalance(req.Amount)

	return c.JSON(http.StatusOK, "your wallet balance has been updated")
}

// Use will use the wallet for given phone number and amount
func Use(c echo.Context) error {
	type request struct {
		Amount      int64  `json:"amount"`
		PhoneNumber string `json:"phoneNumber"`
	}

	req := request{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	wallet, err := GetWallet(req.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid phone number")
	}

	wallet.DecreaseBalance(req.Amount)

	return c.JSON(http.StatusOK, "your wallet balance has been updated")
}
