package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {

	conn := "host=localhost  port=5432 user=postgres password=9539Abdu dbname=user sslmode=disable"

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Db not reachable", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id       SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	return db
}