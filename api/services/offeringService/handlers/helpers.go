package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
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

// will generate a random string for offer code
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// NewOfferCode is a constructor for a new offer code
func NewOfferCode(amount int64, usersCount int64) OfferCode {
	code := randomString(12)

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

// check of the code is valid and doesn't reach its limit
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

// UseCode will check if the code is valid and if it is, it calls an API from wallet service and charge the wallet
func UseCode(code string, phoneNumber string) error {
	for i := 0; i < len(ActiveCodes); i++ {
		//ActiveCodes[i].Lock()
		//defer ActiveCodes[i].Unlock()

		if ActiveCodes[i].Code == code {
			res := checkValidation(&ActiveCodes[i])
			log.Println(ActiveCodes[i])

			if res {
				postBody, _ := json.Marshal(map[string]interface{}{
					"amount":      ActiveCodes[i].Amount,
					"phoneNumber": phoneNumber,
				})
				responseBody := bytes.NewBuffer(postBody)

				resp, err := http.Post("http://localhost:8080/charge", "application/json", responseBody)
				if err != nil {
					log.Fatalf("An Error Occured %v", err)
					return err
				}
				defer resp.Body.Close()

				return nil
			}
		}
	}

	return errors.New("enable to use code")
}
