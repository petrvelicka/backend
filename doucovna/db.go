package doucovna

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DbConnector struct {
	db *sql.DB
}

func NewDbConnector(filename string) DbConnector {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	d := DbConnector{db: db}
	d.createTables()
	return d
}

func (d *DbConnector) Close() {
	d.db.Close()
}

func (d *DbConnector) createTables() {
	offers := `CREATE TABLE IF NOT EXISTS offers (
		Teacher integer NOT NULL,
		Subject integer NOT NULL,
		Description text NOT NULL
	)`

	subjects := `
		CREATE TABLE IF NOT EXISTS subjects (
        Name text NOT NULL UNIQUE
 	)`

	tutors := `CREATE TABLE IF NOT EXISTS tutors (
		email text NOT NULL UNIQUE,
		Name text NOT NULL,
		password text NOT NULL,
		bio text NOT NULL
	)`

	_, err := d.db.Exec(offers)
	if err != nil {
		log.Fatal("Error creating table")
	}
	_, err = d.db.Exec(subjects)
	if err != nil {
		log.Fatal("Error creating table")
	}
	_, err = d.db.Exec(tutors)
	if err != nil {
		log.Fatal("Error creating table")
	}
}
