package database

import (
	"backend/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // DB Driver
)

// Open connection to the datbases
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConnectionDb)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
