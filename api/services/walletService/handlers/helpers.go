package handlers

import (
	"database/sql"
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
}

func (w *Wallet) IncreaseBalance(amount int64) {
	w.Lock()
	w.Balance = w.Balance + amount
	w.Unlock()
}

// GetBalance will return the balance of a wallet by its phone number
func GetBalance(phoneNumber string) (int64, error) {
	db, _ := sql.Open("sqlite3", "./test.db")

	w, err := GetByNumber(db, phoneNumber)
	if err != nil {
		return 0, err
	}

	return w.Balance, nil
}

// GetWallet will return a wallet by its phone number
func GetWallet(phoneNumber string) (Wallet, error) {
	db, _ := sql.Open("sqlite3", "./test.db")

	wallet, err := GetByNumber(db, phoneNumber)
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
	db, _ := sql.Open("sqlite3", "./test.db")

	err := database.CheckNumbers(db, phoneNumber)
	if err != nil {
		return err
	}

	err = database.AddData(db, fullName, phoneNumber, 0)
	if err != nil {
		return err
	}

	log.Println("one record added")

	return nil
}

func GetByNumber(db *sql.DB, phoneNumber string) (Wallet, error) {
	wallet := Wallet{}

	numbers := database.GetALlNumbers(db)

	for _, number := range numbers {
		if number == phoneNumber {
			record, err := db.Query("SELECT FullName, PhoneNumber, Balance FROM users WHERE PhoneNumber = '" + phoneNumber + "'")
			if err != nil {
				return Wallet{}, err
			}

			for record.Next() {
				record.Scan(&wallet.FullName, &wallet.PhoneNumber, &wallet.Balance)
			}
		}
	}

	return wallet, nil
}
