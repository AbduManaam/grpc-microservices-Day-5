package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {

	conn := "host=localhost  user=postgres password=9539Abdu dbname=order sslmode=disable port=5432"

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// FIX 6: orders table was never created — repo.Create() would always fail
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id      SERIAL PRIMARY KEY,
		user_id TEXT NOT NULL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	return db
}