package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"offering-service/database"
)

// Redeem handler for redeem API and gets code and phone number and check if the number and code is valid or not
// if data is valid will use code and call Increment Api from other service
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

// NewCode generates a new code by given stats and add it to ActiveCodes
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

// CodeUsers will return all the users that use code by their phone number
func CodeUsers(c echo.Context) error {
	res, err := database.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cant find all users right now")
	}

	return c.JSON(http.StatusOK, res)
}
