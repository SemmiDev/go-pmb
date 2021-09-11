package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func MySQLConnect() (MySqlDB *sql.DB) {
	DSN := MysqlUser + ":" + MysqlPassword + "@(" + MysqlHost + ":" + MysqlPort + ")/" + MysqlDbname + "?parseTime=true&clientFoundRows=true"
	MySqlDB, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = MySqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	MySqlDB.SetMaxOpenConns(50)
	MySqlDB.SetMaxIdleConns(50)
	MySqlDB.SetConnMaxLifetime(time.Minute * 30)

	return
}
