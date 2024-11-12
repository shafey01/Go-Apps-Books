package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	originalPhoneNumbers = []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
)

func main() {

	db, err := sql.Open("sqlite3", "phone_numbers.db")
	if err != nil {
		log.Fatalf("Faild to initialize db: %v", err)
	}
	defer db.Close()
	// stmt, err := db.Prepare("INSERT phone_number in phone_numbers values(?)")
	sqlCreate := `CREATE TABLE IF NOT EXISTS phone_numbers (
        phone_number Text
        );`
	sqlInsert := `INSERT INTO phone_numbers (phone_number)
                VALUES (?);`
	sqlDelete := `DELETE FROM phone_numbers`

	if _, err := db.Exec(sqlCreate); err != nil {

		log.Fatalf("Faild to create table: %v", err)
	}

	if _, err := db.Exec(sqlDelete); err != nil {

		log.Fatalf("Faild to delete from table: %v", err)
	}

	stmt, err := db.Prepare(sqlInsert)
	if err != nil {
		log.Fatalf("Faild to prepare statment %v", err)
	}

	for _, phone_number := range originalPhoneNumbers {
		if _, err := stmt.Exec(phone_number); err != nil {
			log.Fatalf("Faild to insert %v: ", err)
		}
	}

	fmt.Println("Original:")
	rows, err := db.Query("SELECT phone_number FROM phone_numbers")
	if err != nil {
		log.Fatalf("Failed to query statement: %v", err)
	}
	defer rows.Close()

	var phones []string
	for rows.Next() {
		var phone_number string
		if err := rows.Scan(&phone_number); err != nil {

			log.Fatalf("Failed to scan row: %v", err)
		}

		fmt.Println(phone_number)
		phones = append(phones, phone_number)
	}
}
