package repository

import "database/sql"

type TransactionRepositoryDB struct {
	db *sql.DB // Package for GO to handle database
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{db: db}
}
