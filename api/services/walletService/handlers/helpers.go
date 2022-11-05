package handlers

import (
	"log"
	"sync"
	"wallet-service/database"
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

	database.UpdateWallet(w.PhoneNumber, amount)
}

func (w *Wallet) IncreaseBalance(amount int64) {
	w.Lock()
	w.Balance = w.Balance + amount
	w.Unlock()

	database.UpdateWallet(w.PhoneNumber, w.Balance)
}

// GetBalance will return the balance of a wallet by its phone number
func GetBalance(phoneNumber string) (int64, error) {
	w, err := database.GetByNumber(phoneNumber)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}

// GetWallet will return a wallet by its phone number
func GetWallet(phoneNumber string) (Wallet, error) {
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
