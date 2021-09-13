package sql

import (
	"database/sql"
	"time"
)

func SetupMySql(db *sql.DB) *sql.DB {
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Minute * 30)

	return db
}
