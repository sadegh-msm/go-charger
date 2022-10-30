package walletService

import (
	"challange/database"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

type Wallet struct {
	FullName    string `json:"name" bson:"fullName"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber"`
	Password    string `json:"password" bson:"password"`
	Balance     int64  `json:"balance" bson:"balance"`
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
