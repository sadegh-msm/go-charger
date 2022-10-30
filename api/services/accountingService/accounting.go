package accountingService

import (
	"challange/api/services/walletService"
	"challange/database"
	"log"
)

var PhoneNumbers []string

func newAccount(fullName, phoneNumber, password string) (walletService.Wallet, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return walletService.Wallet{}, err
	}
	defer database.Close(client, ctx, cancel)

	for _, number := range PhoneNumbers {
		if number == phoneNumber {
			break
		} else {
			PhoneNumbers = append(PhoneNumbers, phoneNumber)
			break
		}
	}

	wallet := walletService.Wallet{
		FullName:    fullName,
		Password:    password,
		PhoneNumber: phoneNumber,
		Balance:     0,
	}
	res, err := database.InsertOne(client, ctx, "accounts", "members", wallet)
	if err != nil {
		return walletService.Wallet{}, err
	}
	log.Println("one record added")
	log.Println(res)

	return wallet, nil
}
