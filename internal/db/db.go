package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Snippet string `json:"snippet"`
}

type DB struct {
	pool *sql.DB
}

func New() (*DB, error) {
	db, err := sql.Open("sqlite3", "./snip.sqlite")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DB{pool: db}, nil
}

func (db *DB) InsertItem(item Item) error {
	query := `INSERT INTO snippets (snipp) VALUES  ($1)`
	_, err := db.pool.Exec(query, item.Snippet)
	return err
}

func (db *DB) GetAllItems() ([]Item, error) {
	query := `SELECT snipp FROM snippets`
	rows, err := db.pool.Query(query)
	if err != nil {
		return []Item{}, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Snippet)
		if err != nil {
			return []Item{}, err
		}
		items = append(items, item)
	}
	if rows.Err() != nil {
		return []Item{}, rows.Err()
	}
	return items, nil
}

func (db *DB) Close() {
	err := db.pool.Close()
	if err != nil {
		return
	}
}
