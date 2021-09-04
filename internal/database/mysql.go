package database

import (
	"database/sql"
	"github.com/SemmiDev/go-pmb/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySQL struct{}

func (s *MySQL) ifEmpty(env string, def string) string {
	if env != "" {
		return env
	}
	return def
}

func (s *MySQL) Connect() (*sql.DB, error) {
	MysqlHost := s.ifEmpty(config.MysqlHost, "localhost")
	MysqlPort := s.ifEmpty(config.MysqlPort, "3306")
	MysqlDbname := s.ifEmpty(config.MysqlDbname, "go_pmb")
	MysqlUser := s.ifEmpty(config.MysqlUser, "root")
	MysqlPassword := config.MysqlPassword

	DSN := MysqlUser + ":" + MysqlPassword + "@(" + MysqlHost + ":" + MysqlPort + ")/" + MysqlDbname + "?parseTime=true&clientFoundRows=true"
	MysqlDB, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, err
	}

	err = MysqlDB.Ping()
	if err != nil {
		return nil, err
	}

	MysqlDB.SetMaxIdleConns(10)
	MysqlDB.SetMaxOpenConns(100)
	MysqlDB.SetConnMaxLifetime(time.Hour)

	return MysqlDB, nil
}

func (s *MySQL) Close(mysqlDB *sql.DB) error {
	err := mysqlDB.Ping()
	if err != nil {
		return err
	}
	mysqlDB.Close()

	return nil
}
