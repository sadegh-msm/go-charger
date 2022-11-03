package handlers

import (
	"challange/api/services/walletService/database"
	"log"
	"sync"
)

type Wallet struct {
	FullName    string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
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

func GetBalance(phoneNumber string) (int64, error) {
	w, err := database.GetByNumber(phoneNumber)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}

func GetWallet(phoneNumber string) (Wallet, error) {
	database.InitialMigration()

	wallet, err := database.GetByNumber(phoneNumber)
	if err != nil {
		return Wallet{}, err
	}

	w := Wallet{
		FullName:    wallet.FullName,
		PhoneNumber: wallet.PhoneNumber,
		Balance:     wallet.Balance,
	}

	return w, nil
}

func NewAcc(fullName, phoneNumber string) error {
	database.InitialMigration()

	err := database.CheckNumbers(phoneNumber)
	if err != nil {
		return err
	}

	err = database.AddData(fullName, phoneNumber, 0)
	if err != nil {
		return err
	}

	log.Println("one record added")

	return nil
}
