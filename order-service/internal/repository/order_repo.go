package repository

import "database/sql"

type Repo struct {
	DB *sql.DB
}

func (r *Repo) Create(userID string) string {
	return "order Created"
}