package config

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

type Config struct {
	BindAddr string `json:"bind_addr"`
}

func MustLoad() *Config {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configLocation := path.Join(home, ".junkdb")
	err = os.MkdirAll(configLocation, os.ModePerm) // create path if doesn't exist
	if err != nil {
		panic(err)
	}
	configFilePath := path.Join(configLocation, "config.json") //TODO: make this customizable?
	_, err = os.Stat(configFilePath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return defaultConfig()
	}
	if err != nil {
		panic(err)
	}
	contents, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	var cfg Config
	err = json.Unmarshal(contents, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func defaultConfig() *Config {
	return &Config{
		BindAddr: "127.0.0.1:9000",
	}
}
