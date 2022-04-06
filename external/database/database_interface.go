package database

import "database/sql"

type dbInterface interface {
	Open() error
	GetConnect() *sql.DB
	// GetCtx() context.Context
	// Optimize() error
	Close()
}

type dbImpl struct {
	dbInterface
}

// NewDatabase instancia implementação banco de dados.
func NewDatabase() dbImpl {
	return dbImpl{}
}
