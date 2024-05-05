package repository

import (
	"database/sql"
	"log"
)

const dataSource = "file:data.sqlite?cache=shared&_fk=1"

func NewDb() *sql.DB {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}
