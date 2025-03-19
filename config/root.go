package config

import (
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
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
