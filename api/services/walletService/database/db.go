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
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Wallet{})
}

// AddData adds new data into database
func AddData(name, number string, balance int64) error {
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
	wallet := Wallet{
		PhoneNumber: number,
	}

	db.First(&wallet)

	if wallet.FullName == "" {
		return nil
	}

	return errors.New("wrong or used number")
}

// GetByNumber will return the wallet by given number
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

//func AddData(db *sql.DB, FullName string, PhoneNumber string, Balance int) error {
//	records := `INSERT INTO users(FullName, PhoneNumber, Balance) VALUES (?, ?, ?)`
//	query, err := db.Prepare(records)
//	if err != nil {
//		return err
//	}
//
//	_, err = query.Exec(FullName, PhoneNumber, Balance)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func getALlNumbers(db *sql.DB) (res []string) {
//	record, err := db.Query("SELECT PhoneNumber FROM users")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer record.Close()
//
//	for record.Next() {
//		var PhoneNumber string
//		record.Scan(&PhoneNumber)
//
//		res = append(res, PhoneNumber)
//	}
//
//	return
//}
//
//func CheckNumbers(db *sql.DB, number string) error {
//	res := getALlNumbers(db)
//
//	for _, num := range res {
//		if num == number {
//			return errors.New("number is already used")
//		}
//	}
//
//	return nil
//}
