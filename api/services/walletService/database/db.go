package database

import (
	"database/sql"
	"errors"
	"log"
)

func AddData(db *sql.DB, FullName string, PhoneNumber string, Balance int) error {
	records := `INSERT INTO users(FullName, PhoneNumber, Balance) VALUES (?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		return err
	}

	_, err = query.Exec(FullName, PhoneNumber, Balance)
	if err != nil {
		return err
	}

	return nil
}

func GetALlNumbers(db *sql.DB) (res []string) {
	record, err := db.Query("SELECT PhoneNumber FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer record.Close()

	for record.Next() {
		var PhoneNumber string
		record.Scan(&PhoneNumber)

		res = append(res, PhoneNumber)
	}

	return
}

func CheckNumbers(db *sql.DB, number string) error {
	res := GetALlNumbers(db)

	for _, num := range res {
		if num == number {
			return errors.New("number is already used")
		}
	}

	return nil
}
