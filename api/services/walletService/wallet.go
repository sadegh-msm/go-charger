package walletService

import (
	"challange/database"
	"go.mongodb.org/mongo-driver/bson"
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
}
