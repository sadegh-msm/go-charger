package database

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// Wallet creating schema for orm
type Wallet struct {
	gorm.Model
	FullName    string
	PhoneNumber string
	Balance     int64
}

// InitialMigration initial the database (SQLite)
func InitialMigration() {
	var err error

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Wallet{})
}

// AddData adds new data into database
func AddData(name, number string, balance int64) error {
	InitialMigration()

	wallet := Wallet{
		FullName:    name,
		PhoneNumber: number,
		Balance:     balance,
	}

	result := db.Create(&wallet)
	if result.Error != nil {
		log.Println("unable to add wallet to database")
		return result.Error
	}

	return nil
}

// CheckNumbers will checks for phone numbers so the phone number will be unique
func CheckNumbers(number string) error {
	InitialMigration()

	var wallet Wallet
	db.Find(&wallet, "phone_number = ?", number)

	if wallet.FullName == "" {
		return nil
	}

	return errors.New("wrong or used number")
}

// GetByNumber will return the wallet by given number
func GetByNumber(number string) (Wallet, error) {
	InitialMigration()

	wallet := Wallet{
		PhoneNumber: number,
	}

	db.First(&wallet)

	if wallet.FullName == "" {
		return Wallet{}, errors.New("wrong number")
	}

	return wallet, nil
}

func UpdateWallet(phoneNumber string, balance int64) {
	InitialMigration()

	wallet := Wallet{
		PhoneNumber: phoneNumber,
	}

	db.First(&wallet)

	wallet.Balance = balance

	db.Save(&wallet)
}
