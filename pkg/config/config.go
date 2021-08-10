package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		zap.S().Fatal(err.Error())
		log.Fatalf("config.New: %v", err.Error())
	}
	return &configImpl{}
}
