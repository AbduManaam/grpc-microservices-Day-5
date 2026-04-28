package repository

import "database/sql"

type Repo struct {
	DB *sql.DB
}

func (r *Repo) Create(userID string) (string, error) {
	_, err := r.DB.Exec(
		"INSERT INTO orders (user_id) VALUES ($1)",
		userID,
	)
	if err != nil {
		return "", err
	}
	return "order created", nil
}