package handlers

import (
	"challange/api/services/accountingService"
	"challange/api/services/offeringService"
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

	err := accountingService.NewAccount(req.fullName, req.phoneNumber, req.password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "unable to create account")
	}

	return c.JSON(http.StatusCreated, "account created")
}

func Redeem(c echo.Context) error {
	type request struct {
		phoneNumber string
		code        string
	}

	req := &request{}
	bindErr := c.Bind(req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	err := offeringService.UseCode(req.code, req.phoneNumber)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "invalid code or code has been used too much")
	}

	return c.JSON(http.StatusOK, "your wallet has been charged")
}
