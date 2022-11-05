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
		PhoneNumber string `json:"phoneNumber"`
		Code        string `json:"code"`
	}

	req := request{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	err := UseCode(req.Code, req.PhoneNumber)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "invalid code or code has been used too much")
	}

	err = database.Set(req.PhoneNumber, req.Code)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "you cant use code")
	}

	return c.JSON(http.StatusOK, "your wallet has been charged")
}

// NewCode generates a new code by given stats and add it to ActiveCodes
func NewCode(c echo.Context) error {
	type request struct {
		Amount    int `json:"amount"`
		UserCount int `json:"userCount"`
	}

	req := request{}
	bindErr := c.Bind(&req)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	code := NewOfferCode(int64(req.Amount), int64(req.UserCount))

	return c.JSON(http.StatusOK, code)
}

// CodeUsers will return all the users that use code by their phone number
func CodeUsers(c echo.Context) error {
	var allUsers []string

	res, err := database.GetAll("*")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cant find all users right now")
	}

	for _, item := range res {
		result, err := database.Get(item)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cant find all users right now")
		}

		allUsers = append(allUsers, item+" : "+string(result.([]uint8)))
	}

	return c.JSON(http.StatusOK, allUsers)
}
