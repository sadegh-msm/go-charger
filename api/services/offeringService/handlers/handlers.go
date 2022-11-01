package handlers

import (
	"challange/api/services/offeringService"
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

	err := offeringService.UseCode(req.code, req.phoneNumber)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "invalid code or code has been used too much")
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

	code := offeringService.NewOfferCode(int64(req.amount), int64(req.userCount))

	return c.JSON(http.StatusOK, code)

}
