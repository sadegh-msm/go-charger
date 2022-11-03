package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type OfferCode struct {
	Code      string
	Amount    int64
	UsersCap  int64
	UsedUsers int64
	IsValid   bool
	sync.Mutex
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
	ActiveCodes = append(ActiveCodes, offerCode)

	return offerCode
}

func checkValidation(code *OfferCode) bool {
	if code.IsValid {
		if code.UsedUsers <= code.UsersCap {
			code.Lock()
			code.UsedUsers++
			code.Unlock()

			return true
		} else {
			code.IsValid = false
			return false
		}
	}
	return false
}

func UseCode(code string, phoneNumber string) error {
	for i := 0; i < len(ActiveCodes); i++ {
		if ActiveCodes[i].Code == code {
			res := checkValidation(&ActiveCodes[i])
			if res {
				postBody, _ := json.Marshal(map[string]string{
					"amount":      strconv.FormatInt(ActiveCodes[i].Amount, 10),
					"phoneNumber": phoneNumber,
				})
				responseBody := bytes.NewBuffer(postBody)

				resp, err := http.Post("http://localhost:8080/increment", "application/json", responseBody)
				if err != nil {
					log.Fatalf("An Error Occured %v", err)
				}
				defer resp.Body.Close()
			}
		}
	}

	return nil
}
