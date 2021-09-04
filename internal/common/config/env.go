package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	AppPort           string
	MysqlHost         string
	MysqlPort         string
	MysqlDbname       string
	MysqlUser         string
	MysqlPassword     string
	MidtransServerKey string
	MidtransClientKey string
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

var config = &configImpl{}

func Load(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic(err)
	}

	AppPort = config.Get("")
	AppPort = config.Get("APP_PORT")
	MysqlHost = config.Get("MYSQL_HOST")
	MysqlPort = config.Get("MYSQL_PORT")
	MysqlDbname = config.Get("MYSQL_DBNAME")
	MysqlUser = config.Get("MYSQL_USER")
	MysqlPassword = config.Get("MYSQL_PASSWORD")
	MidtransServerKey = config.Get("MIDTRANS_SERVER_KEY")
	MidtransClientKey = config.Get("MIDTRANS_CLIENT_KEY")
}
