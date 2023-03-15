package config

import (
	"mygram/commons/exceptions"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type ConfigImpl struct {
}

func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exceptions.PanicIfNeeded(err)
	return &ConfigImpl{}
}

func NewConfig() Config {
	return &ConfigImpl{}
}
