package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
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

	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (c *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func (c *configImpl) ifEmpty(env string, defaultValue string) string {
	if env != "" {
		return env
	}
	return defaultValue
}

var c = &configImpl{}

func Load() {
	AppPort = c.ifEmpty(c.Get("AppPort"), ":3000")

	MysqlHost = c.ifEmpty(c.Get("MYSQL_HOST"), "localhost")
	MysqlPort = c.ifEmpty(c.Get("MYSQL_PORT"), "3306")
	MysqlDbname = c.ifEmpty(c.Get("MYSQL_DBNAME"), "go_pmb")
	MysqlUser = c.ifEmpty(c.Get("MYSQL_USER"), "root")
	MysqlPassword = c.Get("MYSQL_PASSWORD")

	MidtransServerKey = c.Get("MIDTRANS_SERVER_KEY")
	MidtransClientKey = c.Get("MIDTRANS_CLIENT_KEY")

	TokenSymmetricKey = c.Get("TOKEN_SYMMETRIC_KEY")

	accessTokenDuration := c.Get("ACCESS_TOKEN_DURATION")
	duration, _ := time.ParseDuration(accessTokenDuration)
	AccessTokenDuration = duration
}
