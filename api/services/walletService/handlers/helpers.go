package handlers

import (
	"challange/database"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"sync"
)

type Wallet struct {
	FullName    string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Balance     int64  `json:"balance"`
	sync.Mutex
}

func (w *Wallet) DecreaseBalance(amount int64) {
	w.Lock()
	w.Balance = w.Balance - amount
	w.Unlock()
}

func (w *Wallet) IncreaseBalance(amount int64) {
	w.Lock()
	w.Balance = w.Balance + amount
	w.Unlock()
}

func GetWallet(phoneNumber string) (Wallet, error) {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return Wallet{}, err
	}
	defer database.Close(client, ctx, cancel)

	option := bson.D{{"_id", 0}}
	filter := bson.D{
		{"phoneNumber", bson.D{{"$eq", phoneNumber}}},
	}

	cursor, err := database.Query(client, ctx, "accounts", "members", filter, option)
	if err != nil {
		panic(err)
	}

	var result Wallet
	if err := cursor.All(ctx, &result); err != nil {
		panic(err)
	}

	return result, nil
}

var PhoneNumbers []string

func NewAcc(fullName, phoneNumber, password string) error {
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
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

	wallet := Wallet{
		FullName:    fullName,
		Password:    password,
		PhoneNumber: phoneNumber,
		Balance:     0,
	}
	res, err := database.InsertOne(client, ctx, "accounts", "members", wallet)
	if err != nil {
		return err
	}
	log.Println("one record added")
	log.Println(res)

	return nil
}
