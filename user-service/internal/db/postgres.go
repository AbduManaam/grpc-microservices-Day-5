package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "9539Abdu"
	}
	dbname := os.Getenv("POSTGRES_DB")
	if dbname == "" {
		dbname = "user"
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn) // sql.Open does NOT actually connect immediately
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil { // It checks if the database is actually reachable and working. Verifies:
		log.Fatal("Db not reachable", err) // DB server is running, credentials are correct, network is reachable
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
