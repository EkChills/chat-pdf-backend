package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() (error) {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		fmt.Println("couldnt open db: ", err)
		return err
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("couldnt ping db")
		return err
	}

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)

	DB = db

	err = createTables()

	if err != nil {
		fmt.Println("couldnt create table:", err)
		return err
	}
	return err
}

func createTables() (error) {
	query := `
		CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
		)
	`

	createPdfTable := `
	CREATE TABLE IF NOT EXISTS pdf (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	pdfText TEXT,
	userId INTEGER 
	FOREIGN KEY (userId) REFERENCES user(id)
	)
	`

	var querySlice []string = []string{query, createPdfTable}  

	for _, item := range querySlice {
		_, err := DB.Exec(item)

		if err != nil {
			fmt.Println("error creating table: ", err)
			return err
		}
		
	}

	return nil

}