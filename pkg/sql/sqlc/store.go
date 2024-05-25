package db

import "database/sql"

type Store struct {
	Querier
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Querier: New(db),
		db:      db,
	}
}
