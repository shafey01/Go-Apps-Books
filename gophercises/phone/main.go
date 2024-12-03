package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

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

	//////////////////////////////////////////////////////////////////
	sqlCreate := `CREATE TABLE IF NOT EXISTS phone_numbers (
        phone_number Text
        );`
	sqlInsert := `INSERT INTO phone_numbers (phone_number)
                VALUES (?);`
	sqlDelete := `DELETE FROM phone_numbers`
	sqlUpdate := `UPDATE phone_numbers SET phone_number = ? WHERE phone_number = ?`
	sqlDeleteRow := `DELETE FROM phone_numbers WHERE phone_number = ?`
	sqlQuery := `SELECT 1 FROM phone_numbers WHERE phone_number = ? LIMIT 1`
	//////////////////////////////////////////////////////////////////

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

	// fmt.Println("Original:")
	rows, err := db.Query("SELECT phone_number FROM phone_numbers")
	if err != nil {
		log.Fatalf("Failed to query statement: %v", err)
	}
	defer rows.Close()

	var phones []string

	var formatedPhones []string
	for rows.Next() {
		var phone_number string
		if err := rows.Scan(&phone_number); err != nil {

			log.Fatalf("Failed to scan row: %v", err)
		}

		// fmt.Println(phone_number)
		phones = append(phones, phone_number)
		formatedPhones = append(formatedPhones, formatPhoneNumber(phone_number))
	}
	if err := rows.Err(); err != nil {

		log.Fatalf("Failed to scan row: %v", err)

	}

	stmtupadate, err := db.Prepare(sqlUpdate)
	if err != nil {
		log.Fatalf("faild to prepare update statment %v", err)
	}

	stmtQuery, err := db.Prepare(sqlQuery)
	if err != nil {
		log.Fatalf("Faild to prepare Query statment %v", err)
	}
	stmtDelete, err := db.Prepare(sqlDeleteRow)
	if err != nil {
		log.Fatalf("Faild to prepare delete statment %v", err)
	}
	for i, p := range phones {
		fp := formatedPhones[i]
		res, err := stmtQuery.Query(fp)
		if err != nil {

			log.Fatalf("Faild to Query record %v", err)
		}

		duplicate := res.Next()
		res.Close()
		if duplicate {
			if _, err := stmtDelete.Exec(p); err != nil {

				log.Fatalf("Faild to Delete record %v", err)
			}
			continue
		}

		if _, err := stmtupadate.Exec(fp, p); err != nil {

			log.Fatalf("Faild to update record %v", err)
		}
	}
	fmt.Println("Formated: ")
	rows, err = db.Query("SELECT phone_number FROM phone_numbers")
	if err != nil {
		log.Fatalf("Failed to query statement: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var phone_number string
		if err := rows.Scan(&phone_number); err != nil {

			log.Fatalf("Failed to scan row: %v", err)
		}

		fmt.Println(phone_number)
	}
	// for _, p := range phones {
	// 	f := formatPhoneNumber(p)
	// 	fmt.Println(f)
	// }
}

func formatPhoneNumber(phone_number string) string {
	return regexp.MustCompile("[^\\d]").ReplaceAllString(phone_number, "")
}
