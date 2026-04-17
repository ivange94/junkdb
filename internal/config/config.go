package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	BindAddr string `json:"bind_addr"`
}

const DefaultBindAddr = "127.0.0.1:21911"
const appDirName = ".junkdb"
const dataFileName = "data"
const configFileName = "config.json"

func MustLoad() *Config {
	configDir, err := Dir()
	if err != nil {
		panic(err)
	}
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		panic(err)
	}
	configFilePath := filepath.Join(configDir, configFileName)
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
	applyDefaults(&cfg)
	return &cfg
}

func defaultConfig() *Config {
	cfg := &Config{}
	applyDefaults(cfg)
	return cfg
}

func applyDefaults(cfg *Config) {
	if cfg.BindAddr == "" {
		cfg.BindAddr = DefaultBindAddr
	}
}

func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home directory: %w", err)
	}
	return filepath.Join(home, appDirName), nil
}

func DataFilePath() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, dataFileName), nil
}
