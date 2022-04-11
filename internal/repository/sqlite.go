package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./forum.db?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS User(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL, 
		surname TEXT NOT NULL
	)
	`); err != nil {
		return err
	}
	return nil
}
