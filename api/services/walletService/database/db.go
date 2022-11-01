package database

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Wallet struct {
	gorm.Model
	FullName    string
	PhoneNumber string
	Password    string
	Balance     int64
}

func InitialMigration() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Wallet{})
}

func AddData(name, number, pass string, balance int64) error {
	wallet := Wallet{
		FullName:    name,
		PhoneNumber: number,
		Password:    pass,
		Balance:     balance,
	}

	result := db.Create(&wallet)
	if result.Error != nil {
		log.Println("unable to add wallet to database")
		return result.Error
	}

	return nil
}

func CheckNumbers(number string) error {
	wallet := Wallet{
		PhoneNumber: number,
	}

	db.First(&wallet)

	if wallet.FullName == "" {
		return nil
	}

	return errors.New("wrong or used number")
}

func GetByNumber(number string) (Wallet, error) {
	wallet := Wallet{
		PhoneNumber: number,
	}

	db.First(&wallet)

	if wallet.FullName == "" {
		return Wallet{}, errors.New("wrong number")
	}

	return wallet, nil
}
