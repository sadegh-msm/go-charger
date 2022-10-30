package offeringService

import (
	"challange/api/services/walletService"
	"errors"
	"math/rand"
)

type OfferCode struct {
	Code      string
	Amount    int64
	UsersCap  int64
	UsedUsers int64
	IsValid   bool
}

var ActiveCodes []OfferCode

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func NewOfferCode(amount int64, usersCount int64) OfferCode {
	code := RandomString(12)
	offerCode := OfferCode{
		Amount:    amount,
		Code:      code,
		UsersCap:  usersCount,
		UsedUsers: 0,
		IsValid:   true,
	}

	return offerCode
}

func checkValidation(code OfferCode, userCode string) bool {
	if code.UsedUsers >= code.UsersCap {
		code.IsValid = false
		return false
	}

	if code.Code != userCode || !code.IsValid {
		return false
	}

	return true
}

func UseCode(code string, phoneNumber string) error {
	var offerCode OfferCode

	for _, item := range ActiveCodes {
		result := checkValidation(item, code)
		if result == true {
			offerCode = item
			break
		}
		return errors.New("code is not valid")
	}

	wallet, err := walletService.GetWallet(phoneNumber)
	if err != nil {
		return errors.New("cant find the wallet")
	}
	wallet.IncreaseBalance(offerCode.Amount)

	return nil
}
