package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SqlDb struct {
	DSN string
}

func NewSqlDb(DSN string) *SqlDb {
	return &SqlDb{DSN: DSN}
}

func (s *SqlDb) Open() *sqlx.DB {
	db, err := sqlx.Open("mysql", s.DSN)
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Minute * 30)

	return db
}
