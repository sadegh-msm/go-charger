package handlers

import (
	"challange/api/services/offeringService/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

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

	err := UseCode(req.code, req.phoneNumber)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "invalid code or code has been used too much")
	}

	err = database.Set(req.phoneNumber, req.code)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "you cant use code")
	}
	return c.JSON(http.StatusOK, "your wallet has been charged")
}

func NewCode(c echo.Context) error {
	type request struct {
		amount    int
		userCount int
	}

	req := &request{}
	bindErr := c.Bind(req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	code := NewOfferCode(int64(req.amount), int64(req.userCount))

	return c.JSON(http.StatusOK, code)
}

func CodeUsers(c echo.Context) error {
	res := database.GetAll()

	return c.JSON(http.StatusOK, res)
}
