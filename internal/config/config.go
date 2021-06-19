package config

import (
	"github.com/joho/godotenv"
	exception2 "go-clean/internal/exception"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception2.PanicIfNeeded(err)
	return &configImpl{}
}