package db

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

func NewDB()*sql.DB{

	conn:= "host=postgres user=postgres password=9539Abdu dbname=order sslmode=disable port=5432"

	db,err:= sql.Open("postgres",conn)
	if err!=nil{
		log.Fatal(err)
	}
	if err:= db.Ping();err!=nil{
		log.Fatal(err)
	}
	return db
}