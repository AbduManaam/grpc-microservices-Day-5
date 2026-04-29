package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {

	conn := "host=localhost  port=5432 user=postgres password=9539Abdu dbname=user sslmode=disable"

	db, err := sql.Open("postgres", conn)    //sql.Open() does NOT actually connect immediately
                                             // It just prepares the connection pool
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {          //It checks if the database is actually reachable and working.Verifies:
       log.Fatal("Db not reachable", err)     //DB server is running,Credentials are correct,Network is reachable
		   
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