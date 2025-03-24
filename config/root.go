package config

import (
	"github.com/joho/godotenv"
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	DB struct {
		Database string
		URL      string
	} `toml:"db"`

	Info struct {
		Port    string
		Service string
	} `toml:"info"`

	Aws struct {
		Key       string
		SecretKey string
		Region    string
		Bucket    string
	} `toml:"aws"`
}

func NewConfig(path string) *Config {
	_ = godotenv.Load(".env")

	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		c.Aws.Key = os.Getenv("AWS_ACCESS_KEY_ID")
		c.Aws.SecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
		return c
	}
}
