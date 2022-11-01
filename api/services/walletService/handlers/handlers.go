package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewAccount(c echo.Context) error {
	type request struct {
		fullName    string
		phoneNumber string
		password    string
	}

	req := &request{}
	bindErr := c.Bind(req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	err := NewAcc(req.fullName, req.phoneNumber, req.password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "unable to create account")
	}

	return c.JSON(http.StatusCreated, "account created")
}
