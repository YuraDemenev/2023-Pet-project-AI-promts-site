package repository

import (
	"fmt"
	site "site/pkg/elements"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user site.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2) RETURNING id", usersTable)
	// row := r.db.QueryRow(query, user.UserName, user.Password)

	// if err := row.Scan(&id); err != nil {
	// 	return 0, err
	// }
	row := tx.QueryRow(query, user.UserName, user.Password)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *AuthPostgres) GetUser(username, password string) (site.User, error) {
	var user site.User
	query_ := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query_, username, password)

	return user, err
}
