package database

import (
	"database/sql"
	"fmt"
	"go-package/configapp"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Client database.
var db *sql.DB

//****************************************************

// GetConnect retorna conexão database.
func (d dbImpl) GetConnect() *sql.DB {
	return db
}

// Open abre database.
func (d dbImpl) Open() error {
	if db != nil {
		return nil
	}

	const (
		dbDriver string = "pgx"
		dbType   string = "postgres"
	)

	cfg := configapp.GetData()

	dbURI := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		dbType,
		cfg.Db.User,
		cfg.Db.Pwd,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name,
	)

	var err error
	db, err = sql.Open(dbDriver, dbURI)
	if err != nil {
		return err
	}
	return nil
}

// Close fecha conexão database.
func (d dbImpl) Close() {
	db.Close()
}
