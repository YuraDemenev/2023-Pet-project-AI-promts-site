package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ImagePostgres struct {
	db *sqlx.DB
}

func NewImagePostgres(db *sqlx.DB) *ImagePostgres {
	return &ImagePostgres{db: db}
}

func (r *ImagePostgres) SavePromts(promts []string, imageUrl string) error {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, str := range promts {
		str = strings.ToLower(str)
		query := fmt.Sprintf("INSERT INTO %s (title, image_url) VALUES ($1, $2)", promtsTable)
		// _, err := r.db.Query(query, str, imageUrl)
		_, err := tx.Exec(query, str, imageUrl)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (r *ImagePostgres) SaveLink(id int, imageLink string) error {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (user_id, image_url,like_count) VALUES ($1, $2, $3)", imagesTable)
	// _, err := r.db.Query(query, id, imageLink, 0)
	_, err = tx.Exec(query, id, imageLink, 0)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
